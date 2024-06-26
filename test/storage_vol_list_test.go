package test

import (
	"fmt"
	"log"
	"net/url"
	"testing"

	"github.com/digitalocean/go-libvirt"
)

func TestVolList(t *testing.T) {

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

	vols, _, err := l.StoragePoolListAllVolumes(p, 1, 0)
	if err != nil {
		log.Printf("failed to retrieve vols: %v", err)
	}

	fmt.Println("Pool\t\tVolName\t\t\t\tVolPath")
	fmt.Printf("----------------------------------------------------------------------------\n")
	for _, v := range vols {
		fmt.Printf("%s\t%s\t\t%s\n", v.Pool, v.Name, v.Key)
	}

	if err = l.Disconnect(); err != nil {
		log.Printf("failed to disconnect: %v", err)
	}
}
