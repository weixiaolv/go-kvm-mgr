package test

import (
	"log"
	"net/url"
	"testing"

	"github.com/digitalocean/go-libvirt"
)

func TestVolDelete(t *testing.T) {

	//uri, _ := url.Parse(string(libvirt.QEMUSystem))
	uri, _ := url.Parse(string("qemu+ssh://root@10.10.54.220/system"))
	l, err := libvirt.ConnectToURI(uri)
	if err != nil {
		log.Printf("failed to connect: %v", err)
	}

	p, err := l.StoragePoolLookupByName("data-sys")
	if err != nil {
		log.Printf("Storage pool lookup failure: %v", err)
		return
	}

	v, err := l.StorageVolLookupByName(p, "test22-sys.qcow2")
	if err != nil {
		log.Printf("failed to retrieve vols: %v", err)
	}

	err = l.StorageVolDelete(v, 0)
	if err != nil {
		log.Printf("storage vol delete failed: %v", err)
	}

	if err = l.Disconnect(); err != nil {
		log.Printf("failed to disconnect: %v", err)
	}
}
