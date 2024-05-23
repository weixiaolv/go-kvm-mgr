package main

import (
	"fmt"
    "github.com/digitalocean/go-libvirt"
    "log"
    "net/url"
)

func main() {
    //uri, _ := url.Parse(string(libvirt.QEMUSystem))
    uri, _ := url.Parse(string("qemu+ssh://root@10.10.54.220/system"))
    l, err := libvirt.ConnectToURI(uri)
    if err != nil {
        log.Printf("failed to connect: %v", err)
    }

    //flags := libvirt.ConnectListStoragePoolsActive | libvirt.ConnectListStoragePoolsInactive

    xmlDesc := `
        <pool type='dir'>
          <name>test</name>
          <target>
            <path>/tmp/test</path>
            <permissions>
              <mode>0755</mode>
              <owner>0</owner>
              <group>0</group>
            </permissions>
          </target>
        </pool>
    `
    pool, err := l.StoragePoolDefineXML(xmlDesc, 0)
    if err != nil {
        log.Printf("Failed to define storage pool: %v", err)
    }

    status, capacity, allocation, available, err := l.StoragePoolGetInfo(pool)

    if err != nil {
        log.Printf("get storage pool info error: %v", err)
    }

    fmt.Printf("status: %v \n capacity: %v \n allocation: %v \n available: %v \n", status, capacity, allocation, available )

    if err = l.Disconnect(); err != nil {
        log.Printf("failed to disconnect: %v", err)
    }
}

