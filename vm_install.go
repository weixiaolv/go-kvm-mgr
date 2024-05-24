package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/weixiaolv/go-kvm-mgr/utils"
)

func main() {
	host := flag.String("host", "", "KVM宿主机IP")
	template := flag.String("template", "template", "vm模板名称")
	name := flag.String("name", "", "新vm名称")
	cpu := flag.Int("cpu", 1, "CPU核心数")
	memory := flag.Int("memory", 512, "内存，单位MB")
	disk := flag.Int("disk", 0, "系统磁盘默认30G，如果需要添加磁盘，请输入该参数，不需要可忽略该参数，单位GB")
	diskMount := flag.String("diskMount", "", "如果添加了disk参数，可添加该参数指明新加磁盘的挂载路径")
	ip := flag.String("ip", "", "新vm IP地址，IP格式ip/subnet，例如192.168.1.10/24")
	gateway := flag.String("gateway", "", "新vm网关，例如：192.168.1.254")
	vlan := flag.String("vlan", "", "虚拟机网络")

	flag.Parse()

	log.Println("kvm_host", *host)
	log.Println("template", *template)
	log.Println("name", *name)
	log.Println("cpu", *cpu)
	log.Println("memory", *memory)
	log.Println("disk", *disk)
	log.Println("diskMount", *diskMount)
	log.Println("ip", *ip)
	log.Println("gateway", *gateway)
	log.Println("vlan", *vlan)

	const (
		// 系统磁盘池
		poolSysName string = "data-sys"
		// 系统磁盘池，卷文件目录
		poolSysDir string = "/data/sys/"
		// 数据磁盘池
		poolDataName string = "data-volumes"
		// 数据磁盘池，卷文件目录
		poolDataDir string = "/data/volumes/"
	)

	volSys := utils.VolNameBuilder(*name, "-sys.qcow2")
	volSysPath := utils.VolNameBuilder(poolSysDir, volSys)
	volData := utils.VolNameBuilder(*name, "-data_1.qcow2")

	conn, _ := utils.GetConn(*host)

	// vmObj 是否已经存在
	vmObj, err := conn.DomainLookupByName(*name)
	if err != nil && strings.Contains(err.Error(), "Domain not found") && vmObj.Name == "" {
		log.Printf("not found vmObj, vmObj is not exist: %v", err)
	}
	if err == nil {
		log.Printf("found vmObj, vmObj is exist: %v", vmObj.Name)
		utils.DisConn(conn)
		return
	}

	// PoolSys磁盘池是否存在
	poolSysObj, err := conn.StoragePoolLookupByName(poolSysName)
	if err != nil {
		log.Printf("Not found storage pool %v: %v", poolSysName, err)
		utils.DisConn(conn)
		return
	} else {
		log.Printf("Found system pool: %v", poolSysObj.Name)
	}

	// PoolData磁盘池是否存在
	poolDataObj, err := conn.StoragePoolLookupByName(poolDataName)
	if err != nil {
		log.Printf("Not found storage pool %v: %v", poolDataName, err)
		utils.DisConn(conn)
		return
	} else {
		log.Printf("Found data pool: %v", poolDataObj.Name)
	}

	// sys磁盘文件是否存在
	volSysObj, err := conn.StorageVolLookupByName(poolSysObj, volSys)
	if err == nil {
		log.Printf("volumes is exist: %v", volSysObj.Name)
		utils.DisConn(conn)
		return
	} else {
		log.Printf("%v", err)
	}

	// data磁盘文件是否存在
	volDataObj, err := conn.StorageVolLookupByName(poolDataObj, volData)
	if err == nil {
		log.Printf("volumes is exist: %v", volDataObj.Name)
		utils.DisConn(conn)
		return
	} else {
		log.Printf("%v", err)
	}

	/*
		//dump template xml
		templateObj, err := conn.DomainLookupByName(*template)
		if err != nil {
		   log.Printf("template vm not found: %v", err)
		   utils.DisConn(conn)
		   return
		}
		templateXml, err := conn.DomainGetXMLDesc(templateObj, 0)
		if err == nil {
		   log.Printf("%v\n", templateXml)
		   log.Printf("Type of templateXml %T", templateXml)
		}
	*/

	// clone vm volumes
	volXmlDesc := `
       <volume type='file'>
         <name>targetVolName</name>
         <key>targetVolPath</key>
         <source>
         </source>
         <target>
           <path>targetVolPath</path>
           <format type='qcow2'/>
           <permissions>
             <mode>0644</mode>
             <owner>64055</owner>
             <group>64055</group>
           </permissions>
         </target>
       </volume>
    `
	targetVolXmlDesc1 := strings.ReplaceAll(volXmlDesc, "targetVolName", volSys)
	targetVolXmlDesc2 := strings.ReplaceAll(targetVolXmlDesc1, "targetVolPath", volSysPath)

	templateVolName := utils.VolNameBuilder(*template, ".qcow2")
	templateVolObj, err := conn.StorageVolLookupByName(poolSysObj, templateVolName)
	if err != nil {
		log.Printf("Failed to find template volume: %v", err)
		utils.DisConn(conn)
		return
	}

	volSysObj, err = conn.StorageVolCreateXMLFrom(poolSysObj, targetVolXmlDesc2, templateVolObj, 0)
	if err != nil {
		log.Printf("Failed to create storage volume: %v", err)
		utils.DisConn(conn)
		return
	} else {
		log.Printf("create storage volume success %v", volSysObj.Name)
	}

	// 修改替换虚拟机xml文件对应参数
	vmXmlDesc, err := utils.VMXmlHandle(*name, *cpu, *memory, *vlan, volSysPath)
	if err != nil {
		log.Printf("VMXmlHandle error: %v", err)
		utils.DisConn(conn)
		return
	}
	log.Println(vmXmlDesc)

	// 创建虚拟机(define)
	newVM, err := conn.DomainDefineXMLFlags(vmXmlDesc, 1)
	if err != nil {
		log.Printf("VM define error: %v", err)
		utils.DisConn(conn)
		return
	}

	// 启动虚拟机(start)
	err = conn.DomainCreate(newVM)
	if err != nil {
		log.Printf("VM start failed: %v", err)
		utils.DisConn(conn)
		return
	}

	// 查询虚拟机启动成功状态
	log.Println("等待5s虚拟机启动...")
	time.Sleep(5 * time.Second)
	log.Println("开始查询虚拟机启动状态")

	timeout := 10 * time.Second
	deadLine := time.Now().Add(timeout)
	for now := range time.Tick(2 * time.Second) {
		if now.After(deadLine) {
			log.Println("查询虚拟机启动状态超时(10秒)")
			utils.DisConn(conn)
			return
		}

		isActive, err := conn.DomainIsActive(newVM)
		if err != nil {
			log.Printf("虚拟机启动失败,继续查询状态: %v", err)
			continue
		}
		if isActive == 0 {
			log.Printf("虚拟机处于 inactive 状态: %v", err)
			continue
		}
		if isActive == 1 {
			log.Printf("虚拟机启动成功")
			break
		}

	}

	// 等待qemu-agent 启动
	time.Sleep(5 * time.Second)

	// 修改 vm 主机名
	hostname := fmt.Sprintf(`{
            "execute":"guest-exec",
            "arguments":{
                "path":"/usr/bin/hostnamectl",
                "arg":["set-hostname", "%s"]
                }
            }`, *name)
	_, err = conn.QEMUDomainAgentCommand(newVM, hostname, 3, 0)
	if err != nil {
		log.Printf("修改 vm 主机名失败:%v", err)
	} else {
		log.Printf("修改主机名完成")
	}

	// 修改 vm hosts
	hosts := fmt.Sprintf(`{
            "execute":"guest-exec",
            "arguments":{
                "path":"/usr/bin/echo",
                "arg":["%s", "%s", ">>/etc/hosts"]
                }
            }`, *ip, *name)
	_, err = conn.QEMUDomainAgentCommand(newVM, hosts, 3, 0)
	if err != nil {
		log.Printf("修改 vm hosts 失败:%v", err)
	}

	//

	utils.DisConn(conn)

}
