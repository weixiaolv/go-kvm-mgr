package test

import (
	"fmt"
	"log"
	"net/url"
	"testing"

	"github.com/digitalocean/go-libvirt"
)

func TestGetXmlDesc(t *testing.T) {

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

	xml, err := l.NetworkGetXMLDesc(net, 0)
	if err != nil {
		log.Printf("failed to get network xml: %v", err)
		return
	}

	fmt.Printf(xml)

	if err = l.Disconnect(); err != nil {
		log.Printf("failed to disconnect: %v", err)
	}
}
