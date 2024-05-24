package test

import (
	"log"
	"net/url"
	"testing"

	"github.com/digitalocean/go-libvirt"
)

func TestVolClone(t *testing.T) {

	uri, _ := url.Parse(string("qemu+ssh://root@10.10.54.220/system"))
	l, err := libvirt.ConnectToURI(uri)
	if err != nil {
		log.Printf("failed to connect: %v", err)
	}

	p, err := l.StoragePoolLookupByName("data-sys")
	if err != nil {
		log.Printf("Storage pool lookup failure: %v", err)
		return
	}

	baseVol, err := l.StorageVolLookupByName(p, "template.qcow2")
	if err != nil {
		log.Printf("Failed to find base volume: %v", err)
		return
	}

	xmlDesc := `
        <volume type='file'>
          <name>test22-sys.qcow2</name>
          <key>/data/sys/test22-sys.qcow2</key>
          <source>
          </source>
          <target>
            <path>/data/sys/test22-sys.qcow2</path>
            <format type='qcow2'/>
            <permissions>
              <mode>0644</mode>
              <owner>64055</owner>
              <group>64055</group>
            </permissions>
          </target>
        </volume>
    `

	_, err = l.StorageVolCreateXMLFrom(p, xmlDesc, baseVol, 0)
	if err != nil {
		log.Printf("Failed to create storage volume: %v", err)
		return
	}

	// 在这里可以使用 vol 进行后续操作
	log.Printf("Storage volume created successfully")
}
