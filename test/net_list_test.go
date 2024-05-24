package test

import (
	"fmt"
	"log"
	"net/url"
	"testing"

	"github.com/digitalocean/go-libvirt"
)

func TestNetList(t *testing.T) {
	//uri, _ := url.Parse(string(libvirt.QEMUSystem))
	uri, _ := url.Parse(string("qemu+ssh://root@10.10.54.220/system"))
	l, err := libvirt.ConnectToURI(uri)
	if err != nil {
		log.Printf("failed to connect: %v", err)
	}

	flags := libvirt.ConnectListNetworksInactive | libvirt.ConnectListNetworksActive |
		libvirt.ConnectListNetworksPersistent | libvirt.ConnectListNetworksTransient |
		libvirt.ConnectListNetworksAutostart | libvirt.ConnectListNetworksNoAutostart

	nets, _, err := l.ConnectListAllNetworks(1, flags)
	if err != nil {
		log.Printf("failed to retrieve nets: %v", err)
	}

	fmt.Println("UUID\t\t\t\t\tName")
	fmt.Printf("--------------------------------------------------------\n")
	for _, n := range nets {
		fmt.Printf("%x\t%s\n", n.UUID, n.Name)
	}

	if err = l.Disconnect(); err != nil {
		log.Printf("failed to disconnect: %v", err)
	}
}
