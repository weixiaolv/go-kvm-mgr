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

    net, err := l.NetworkLookupByName("testvlan151")
    if err != nil {
        log.Printf("failed to find network: %v", err)
        return
    }

    // get bridgename
    a, err := l.NetworkGetBridgeName(net)
    if err != nil {
        log.Printf("failed to get autostart network: %v", err)
    }
    fmt.Printf("bridgename: %v\n", a)

    // get autostart
    b, err := l.NetworkGetAutostart(net)
    if err != nil {
        log.Printf("failed to get autostart network: %v", err)
    }
    fmt.Printf("autostart: %v\n", b)

    // get ports
    ports, _, err := l.NetworkListAllPorts(net, 1, 0)
    if err != nil {
        log.Printf("failed to get autostart network: %v", err)
    }
    for _, p := range ports {
        fmt.Printf("%v, %v", p.UUID, p.Net)
    }




    if err = l.Disconnect(); err != nil {
        log.Printf("failed to disconnect: %v", err)
    }
}

