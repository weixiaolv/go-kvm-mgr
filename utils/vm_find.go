package utils

import (
	"log"
	"strings"

	"github.com/digitalocean/go-libvirt"
)

func VMFind(VMName string, conn *libvirt.Libvirt) int {
	count := 0
	vm, err := conn.DomainLookupByName(VMName)
	if err != nil && strings.Contains(err.Error(), "Domain not found") && vm.Name == "" {
		count = 0
		log.Printf("Faied to find vm: %v", err)
	}
	if err == nil {
		count = 1
		log.Printf("find vm: %v", vm.Name)
	}
	return count
}
