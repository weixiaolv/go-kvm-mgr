package main

import (
	"fmt"
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

    flags := libvirt.ConnectListStoragePoolsActive | libvirt.ConnectListStoragePoolsInactive

    pools, _, err := l.ConnectListAllStoragePools(1, flags)
    if err != nil {
        log.Printf("failed to retrieve pools: %v", err)
    }

    fmt.Println("UUID\t\t\t\t\tName")
    fmt.Printf("--------------------------------------------------------\n")
    for _, p := range pools {
        fmt.Printf("%x\t%s\n", p.UUID, p.Name)
    }

    if err = l.Disconnect(); err != nil {
        log.Printf("failed to disconnect: %v", err)
    }
}

