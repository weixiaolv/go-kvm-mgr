package main

import (
	"log"
    "net/url"

    "github.com/digitalocean/go-libvirt"
)

func main() {
    //uri, _ := url.Parse(string(libvirt.QEMUSystem))
    uri, _ := url.Parse(string("qemu+ssh://root@10.10.54.220/system"))
    l, err := libvirt.ConnectToURI(uri)
    if err != nil {
        log.Printf("failed to connect: %v", err)
    }

    p, err := l.StoragePoolLookupByName("test")
    if err != nil {
        log.Printf("Storage pool lookup failure: %v", err)
        return
    }

    err = l.StoragePoolCreate(p, 0)
    if err == nil {
        log.Printf("Storage pool active success: %v", p.Name)
    }

    if err != nil {
        log.Printf("Storage pool active error: %v", err)
    }

    if err = l.Disconnect(); err != nil {
        log.Printf("failed to disconnect: %v", err)
    }
}

