package test

import (
	"log"
	"net/url"
	"testing"

	"github.com/digitalocean/go-libvirt"
)

func TestNetSetAuto(t *testing.T) {

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

	err = l.NetworkSetAutostart(net, 1)
	if err != nil {
		log.Printf("failed to set autostart network: %v", err)
	}

	if err = l.Disconnect(); err != nil {
		log.Printf("failed to disconnect: %v", err)
	}
}
