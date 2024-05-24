package test

import (
	"fmt"
	"log"
	"net/url"
	"testing"

	"github.com/digitalocean/go-libvirt"
)

func TestVolInfo(t *testing.T) {

	//uri, _ := url.Parse(string(libvirt.QEMUSystem))
	uri, _ := url.Parse(string("qemu+ssh://root@10.10.54.220/system"))
	l, err := libvirt.ConnectToURI(uri)
	if err != nil {
		log.Printf("failed to connect: %v", err)
	}

	p, err := l.StoragePoolLookupByName("data1-sys")
	if err != nil {
		log.Printf("Storage pool lookup failure: %v", err)
		return
	}

	v, err := l.StorageVolLookupByName(p, "test11-sys.qcow2")
	if err != nil {
		log.Printf("failed to retrieve vols: %v", err)
	}

	_, capacity, allocation, err := l.StorageVolGetInfo(v)
	if err != nil {
		log.Printf("failed to get vol info: %v", err)
	} else {
		fmt.Println("PoolName\tVolName\t\t\tVolPath\t\t\t\tVolCapacity(MB)\t\tVolAllocation(MB)")
		fmt.Println("-----------------------------------------------------------------------------------------------------")
		fmt.Printf("%v\t%v\t%v\t%v\t\t\t%v\n", v.Pool, v.Name, v.Key, capacity/1024/1024, allocation/1024/1024)
	}

	if err = l.Disconnect(); err != nil {
		log.Printf("failed to disconnect: %v", err)
	}
}
