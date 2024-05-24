package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/digitalocean/go-libvirt"
	"github.com/weixiaolv/go-kvm-mgr/utils"
)

func TestDomainList(t *testing.T) {

	l, err := utils.GetConn("10.10.54.220")

	v, err := l.ConnectGetLibVersion()
	if err != nil {
		log.Printf("failed to retrieve libvirt version: %v", err)
	}
	fmt.Println("Version:", v)

	flags := libvirt.ConnectListDomainsActive | libvirt.ConnectListDomainsInactive
	//flags := libvirt.ConnectListDomainsActive
	//flags := libvirt.ConnectListDomainsInactive

	domains, _, err := l.ConnectListAllDomains(1, flags)
	if err != nil {
		log.Printf("failed to retrieve domains: %v", err)
	}

	fmt.Println("ID\tName\t\tUUID")
	fmt.Printf("--------------------------------------------------------\n")
	for _, d := range domains {
		fmt.Printf("%d\t%s\t%x\n", d.ID, d.Name, d.UUID)
	}

	if err = l.Disconnect(); err != nil {
		log.Printf("failed to disconnect: %v", err)
	}
}
