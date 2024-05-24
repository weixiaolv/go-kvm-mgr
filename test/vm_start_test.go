package test

import (
	"log"
	"net/url"
	"testing"

	"github.com/digitalocean/go-libvirt"
)

func TestDomainStart(t *testing.T) {

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

	err = l.DomainCreate(d)
	if err == nil {
		log.Printf("Domain start success: %v", d.Name)
	}

	if err != nil {
		log.Printf("Domain start failed: %v", err)
	}

	if err = l.Disconnect(); err != nil {
		log.Printf("failed to disconnect: %v", err)
	}
}
