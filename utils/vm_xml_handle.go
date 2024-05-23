package libs

import (
    "encoding/xml"
    "log"
)

type Domain struct {
    XMLName        xml.Name  `xml:"domain"`
    Type           string    `xml:"type,attr"`
    Name           string    `xml:"name"`
    Memory         Memory    `xml:"memory"`
    CurrentMemory  Memory    `xml:"currentMemory"`
    VCPU           VCPU      `xml:"vcpu"`
    OS             OS        `xml:"os"`
    Features       Features  `xml:"features"`
    CPU            CPU       `xml:"cpu"`
    Clock          Clock     `xml:"clock"`
    OnPowerOff     string    `xml:"on_poweroff"`
    OnReboot       string    `xml:"on_reboot"`
    OnCrash        string    `xml:"on_crash"`
    PM             PM        `xml:"pm"`
    Devices        Devices   `xml:"devices"`
}

type Memory struct {
    Unit   string  `xml:"unit,attr"`
    Value  int     `xml:",chardata"`
}

type VCPU struct {
    Placement  string  `xml:"placement,attr"`
    Value      int     `xml:",chardata"`
}

type OS struct {
    Type  Type  `xml:"type"`
    Boot  Boot  `xml:"boot"`
}

type Type struct {
    Arch     string  `xml:"arch,attr"`
    Machine  string  `xml:"machine,attr"`
    Value    string  `xml:",chardata"`
}

type Boot struct {
    Dev  string  `xml:"dev,attr"`
}

type Features struct {
    ACPI  string  `xml:"acpi"`
    APIC  string  `xml:"apic"`
}

type CPU struct {
    Mode   string  `xml:"mode,attr"`
    Check  string  `xml:"check,attr"`
}

type Clock struct {
    Offset  string   `xml:"offset,attr"`
    Timers  []Timer  `xml:"timer"`
}

type Timer struct {
    Name        string   `xml:"name,attr"`
    TickPolicy  *string  `xml:"tickpolicy,attr,omitempty"`
    Present     *string  `xml:"present,attr,omitempty"`
}

type PM struct {
    SuspendToMem   SuspendState  `xml:"suspend-to-mem"`
    SuspendToDisk  SuspendState  `xml:"suspend-to-disk"`
}

type SuspendState struct {
    Enabled  string  `xml:"enabled,attr"`
}

type Devices struct {
    Emulator     string        `xml:"emulator"`
    Disks        []Disk        `xml:"disk"`
    Controllers  []Controller  `xml:"controller"`
    Interfaces   []Interface   `xml:"interface"`
    Serial       Serial        `xml:"serial"`
    Console      Console       `xml:"console"`
    Channel      Channel       `xml:"channel"`
    Input        []Input       `xml:"input"`
    Graphics     Graphics      `xml:"graphics"`
    Video        Video         `xml:"video"`
    MemBalloon   MemBalloon    `xml:"memballoon"`
}

type Disk struct {
    Type      string         `xml:"type,attr"`
    Device    string         `xml:"device,attr"`
    Driver    Driver         `xml:"driver"`
    Source    *DiskSource    `xml:"source,omitempty"`
    Target    Target         `xml:"target"`
    Readonly  *string        `xml:"readonly,omitempty"`
    Address   *DiskAddress   `xml:"address,omitempty"`
}

type Driver struct {
    Name   string   `xml:"name,attr"`
    Type   string   `xml:"type,attr"`
    Cache  *string  `xml:"cache,attr,omitempty"`
}

type DiskSource struct {
    File  string  `xml:"file,attr"`
}

type Target struct {
    Dev  string  `xml:"dev,attr"`
    Bus  string  `xml:"bus,attr"`
}

type DiskAddress struct {
    Type        string  `xml:"type,attr"`
    Controller  string  `xml:"controller,attr"`
    Bus         string  `xml:"bus,attr"`
    Target      string  `xml:"target,attr"`
    Unit        string  `xml:"unit,attr"`
}

type Controller struct {
    Type     string              `xml:"type,attr"`
    Index    string              `xml:"index,attr"`
    Model    *string             `xml:"model,attr,omitempty"`
    Address  *ControllerAddress  `xml:"address,omitempty"`
    Master   *Master             `xml:"master,omitempty"`
}

type ControllerAddress struct {
    Type           string   `xml:"type,attr"`
    Domain         string   `xml:"domain,attr,omitempty"`
    Bus            string   `xml:"bus,attr,omitempty"`
    Slot           string   `xml:"slot,attr,omitempty"`
    Function       string   `xml:"function,attr,omitempty"`
    Multifunction  *string  `xml:"multifunction,attr,omitempty"`
}

type Master struct {
    StartPort  string  `xml:"startport,attr,omitempty"`
}

type Interface struct {
    Type     string            `xml:"type,attr"`
    Source   InterfaceSource   `xml:"source"`
    Model    InterfaceModel    `xml:"model"`
    //Ip       InterfaceIp       `xml:"ip"`
    //Route    InterfaceRoute    `xml:"route"`
    //Address  InterfaceAddress  `xml:"address"`
}

type InterfaceSource struct {
    Network  string  `xml:"network,attr"`
}

type InterfaceModel struct {
    Type  string  `xml:"type,attr"`
}

//type InterfaceIp struct {
//    Address    string              `xml:"address,attr"`
//    Prefix     string              `xml:"prefix,attr"`
//}

//type InterfaceRoute struct {
//    Address    string    `xml:"address,attr"`
//    Gateway    string    `xml:"gateway,attr"`
//}


//type InterfaceAddress struct {
//    Type      string  `xml:"type,attr,omitempty"`
//    Domain    string  `xml:"domain,attr,omitempty"`
//    Bus       string  `xml:"bus,attr,omitempty"`
//    Slot      string  `xml:"slot,attr,omitempty"`
//    Function  string  `xml:"function,attr,omitempty"`
//}

type Serial struct {
    Type    string        `xml:"type,attr"`
    Target  SerialTarget  `xml:"target"`
}

type SerialTarget struct {
    Type   string             `xml:"type,attr"`
    Port   string             `xml:"port,attr"`
    Model  SerialTargetModel  `xml:"model"`
}

type SerialTargetModel struct {
    Name  string  `xml:"name,attr"`
}

type Console struct {
    Type    string         `xml:"type,attr"`
    Target  ConsoleTarget  `xml:"target"`
}

type ConsoleTarget struct {
    Type  string  `xml:"type,attr"`
    Port  string  `xml:"port,attr"`
}

type Channel struct {
    Type     string          `xml:"type,attr"`
    Source   ChannelSource   `xml:"source"`
    Target   ChannelTarget   `xml:"target"`
    Address  ChannelAddress  `xml:"address"`
}

type ChannelSource struct {
    Mode  string  `xml:"mode,attr"`
    Path  string  `xml:"path,attr"`
}

type ChannelTarget struct {
    Type  string  `xml:"type,attr"`
    Name  string  `xml:"name,attr"`
}

type ChannelAddress struct {
    Type        string  `xml:"type,attr"`
    Controller  string  `xml:"controller,attr"`
    Bus         string  `xml:"bus,attr"`
    Port        string  `xml:"port,attr"`
}

type Input struct {
    Type     string         `xml:"type,attr"`
    Bus      string         `xml:"bus,attr"`
    Address  *InputAddress  `xml:"address,omitempty"`
}

type InputAddress struct {
    Type  string  `xml:"type,attr"`
    Bus   string  `xml:"bus,attr"`
    Port  string  `xml:"port,attr"`
}

type Graphics struct {
    Type            string          `xml:"type,attr"`
    Port            string          `xml:"port,attr"`
    AutoPort        string          `xml:"autoport,attr"`
    Listen          string          `xml:"listen,attr"`
    GraphicsListen  GraphicsListen  `xml:"listen"`
}

type GraphicsListen struct {
    Type     string  `xml:"type,attr"`
    Address  string  `xml:"address,attr"`
}

type Video struct {
    Model    VideoModel    `xml:"model"`
    Address  VideoAddress  `xml:"address"`
}

type VideoModel struct {
    Type     string   `xml:"type,attr"`
    VRam     string   `xml:"vram,attr"`
    Heads    string   `xml:"heads,attr"`
    Primary  string   `xml:"primary,attr"`
}

type VideoAddress struct {
    Type      string  `xml:"type,attr"`
    Domain    string  `xml:"domain,attr"`
    Bus       string  `xml:"bus,attr"`
    Slot      string  `xml:"slot,attr"`
    Function  string  `xml:"function,attr"`
}

type MemBalloon struct {
    Model    string             `xml:"model,attr"`
    Address  MemBalloonAddress  `xml:"address"`
}

type MemBalloonAddress struct {
    Type      string  `xml:"type,attr"`
    Domain    string  `xml:"domain,attr"`
    Bus       string  `xml:"bus,attr"`
    Slot      string  `xml:"slot,attr"`
    Function  string  `xml:"function,attr"`
}

func VMXmlHandle(vmName string, cpuCore int, memory int,
    vlan string, volSysPath string) (vmXmlDesc string, err error) {

    templateXmlDesc := `
        <domain type='kvm'>
          <name>template</name>
          <memory unit='MiB'>1024</memory>
          <currentMemory unit='MiB'>1024</currentMemory>
          <vcpu placement='static'>2</vcpu>
          <os>
            <type arch='x86_64' machine='pc-i440fx-5.2'>hvm</type>
            <boot dev='hd'/>
          </os>
          <features>
            <acpi/>
            <apic/>
          </features>
          <cpu mode='host-model' check='partial'/>
          <clock offset='utc'>
            <timer name='rtc' tickpolicy='catchup'/>
            <timer name='pit' tickpolicy='delay'/>
            <timer name='hpet' present='no'/>
          </clock>
          <on_poweroff>destroy</on_poweroff>
          <on_reboot>restart</on_reboot>
          <on_crash>destroy</on_crash>
          <pm>
            <suspend-to-mem enabled='no'/>
            <suspend-to-disk enabled='no'/>
          </pm>
          <devices>
            <emulator>/usr/bin/qemu-system-x86_64</emulator>
            <disk type='file' device='disk'>
              <driver name='qemu' type='qcow2' cache='none'/>
              <source file='/data/sys/template.qcow2'/>
              <target dev='vda' bus='virtio'/>
            </disk>
            <disk type='file' device='cdrom'>
              <driver name='qemu' type='raw'/>
              <target dev='hda' bus='ide'/>
              <readonly/>
              <address type='drive' controller='0' bus='0' target='0' unit='0'/>
            </disk>
            <controller type='usb' index='0' model='ich9-ehci1'>
              <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x7'/>
            </controller>
            <controller type='usb' index='0' model='ich9-uhci1'>
              <master startport='0'/>
              <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x0' multifunction='on'/>
            </controller>
            <controller type='usb' index='0' model='ich9-uhci2'>
              <master startport='2'/>
              <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x1'/>
            </controller>
            <controller type='usb' index='0' model='ich9-uhci3'>
              <master startport='4'/>
              <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x2'/>
            </controller>
            <controller type='pci' index='0' model='pci-root'/>
            <controller type='ide' index='0'>
              <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x1'/>
            </controller>
            <controller type='virtio-serial' index='0'>
              <address type='pci' domain='0x0000' bus='0x00' slot='0x07' function='0x0'/>
            </controller>
            <interface type='network'>
              <source network='vlan54'/>
              <model type='virtio'/>
            </interface>
            <serial type='pty'>
              <target type='isa-serial' port='0'>
                <model name='isa-serial'/>
              </target>
            </serial>
            <console type='pty'>
              <target type='serial' port='0'/>
            </console>
            <channel type='unix'>
              <source mode='bind' path='/var/lib/libvirt/qemu/org.qemu.guest_agent.0'/>
              <target type='virtio' name='org.qemu.guest_agent.0'/>
              <address type='virtio-serial' controller='0' bus='0' port='1'/>
            </channel>
            <input type='tablet' bus='usb'>
              <address type='usb' bus='0' port='1'/>
            </input>
            <input type='mouse' bus='ps2'/>
            <input type='keyboard' bus='ps2'/>
            <graphics type='vnc' port='-1' autoport='yes' listen='0.0.0.0'>
              <listen type='address' address='0.0.0.0'/>
            </graphics>
            <video>
              <model type='vga' vram='16384' heads='1' primary='yes'/>
              <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x0'/>
            </video>
            <memballoon model='virtio'>
              <address type='pci' domain='0x0000' bus='0x00' slot='0x06' function='0x0'/>
            </memballoon>
          </devices>
        </domain>
    `

    domain := Domain{}
    err = xml.Unmarshal([]byte(templateXmlDesc), &domain)
    if err != nil {
        log.Printf("Error unmarshalling XML: %v", err)
        return "", err
    }

    domain.Name = vmName
    domain.Memory.Value = memory
    domain.CurrentMemory.Value = memory
    domain.VCPU.Value = cpuCore
    domain.Devices.Interfaces[0].Source.Network = vlan
    domain.Devices.Disks[0].Source.File = volSysPath

    xmlBytes, err := xml.MarshalIndent(domain, "", "  ")
    if err != nil {
        log.Println("Error marshalling XML:", err)
        return "", err
    }

    modifiedXmlString := string(xmlBytes)
    return modifiedXmlString, nil
}
