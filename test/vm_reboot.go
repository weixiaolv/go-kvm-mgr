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

    d, err := l.DomainLookupByName("test11")
    if err != nil {
        log.Printf("Domain lookup failure: %v", err)
        return
    }

    err = l.DomainReboot(d, 0)
    if err == nil {
        log.Printf("Domain reboot success: %v", d.Name)
    }

    if err != nil {
        log.Printf("Domain reboot failed: %v", err)
    }

    if err = l.Disconnect(); err != nil {
        log.Printf("failed to disconnect: %v", err)
    }
}

