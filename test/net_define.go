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

    xmlDesc := `
        <network>
          <name>testvlan151</name>
          <forward mode='bridge'/>
          <bridge name='br151'/>
        </network>
    `
    net, err := l.NetworkDefineXML(xmlDesc)
    if err != nil {
        log.Printf("Failed to define network: %v", err)
    }

    xml,err := l.NetworkGetXMLDesc(net, 0)
    if err != nil {
        log.Printf("failed to get network xml: %v", err)
    }

    fmt.Printf(xml)

    if err = l.Disconnect(); err != nil {
        log.Printf("failed to disconnect: %v", err)
    }
}

