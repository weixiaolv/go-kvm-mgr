package test

import (
	"log"
	"net/url"
	"testing"

	"github.com/digitalocean/go-libvirt"
)

func TestNetStart(t *testing.T) {
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

	err = l.NetworkCreate(net)
	if err != nil {
		log.Printf("failed to start network: %v", err)
	}

	if err = l.Disconnect(); err != nil {
		log.Printf("failed to disconnect: %v", err)
	}
}
