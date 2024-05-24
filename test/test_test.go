package test

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

type Domain struct {
	XMLName       xml.Name `xml:"domain"`
	Type          string   `xml:"type,attr"`
	Name          string   `xml:"name"`
	UUID          string   `xml:"uuid"`
	Memory        Memory   `xml:"memory"`
	CurrentMemory Memory   `xml:"currentMemory"`
	Vcpu          Vcpu     `xml:"vcpu"`
	OS            OS       `xml:"os"`
	Features      Features `xml:"features"`
	CPU           CPU      `xml:"cpu"`
	Clock         Clock    `xml:"clock"`
	OnPoweroff    string   `xml:"on_poweroff"`
	OnReboot      string   `xml:"on_reboot"`
	OnCrash       string   `xml:"on_crash"`
	PM            PM       `xml:"pm"`
	Devices       Devices  `xml:"devices"`
}

type Memory struct {
	Unit  string `xml:"unit,attr"`
	Value int    `xml:",chardata"`
}

type Vcpu struct {
	Placement string `xml:"placement,attr"`
	Value     int    `xml:",chardata"`
}

type OS struct {
	Type Type `xml:"type"`
	Boot Boot `xml:"boot"`
}

type Type struct {
	Arch    string `xml:"arch,attr"`
	Machine string `xml:"machine,attr"`
	Value   string `xml:",chardata"`
}

type Boot struct {
	Dev string `xml:"dev,attr"`
}

type Features struct {
	ACPI string `xml:"acpi"`
	APIC string `xml:"apic"`
}

type CPU struct {
	Mode  string `xml:"mode,attr"`
	Check string `xml:"check,attr"`
}

type Clock struct {
	Offset string  `xml:"offset,attr"`
	Timer  []Timer `xml:"timer"`
}

type Timer struct {
	Name       string `xml:"name,attr"`
	TickPolicy string `xml:"tickpolicy,attr"`
	Present    string `xml:"present,attr,omitempty"`
}

type PM struct {
	SuspendToMem  SuspendTo `xml:"suspend-to-mem"`
	SuspendToDisk SuspendTo `xml:"suspend-to-disk"`
}

type SuspendTo struct {
	Enabled string `xml:"enabled,attr"`
}

type Devices struct {
	Emulator   string       `xml:"emulator"`
	Disk       []Disk       `xml:"disk"`
	Controller []Controller `xml:"controller"`
	Interface  Interface    `xml:"interface"`
	Serial     Serial       `xml:"serial"`
	Console    Console      `xml:"console"`
	Channel    Channel      `xml:"channel"`
	Input      []Input      `xml:"input"`
	Graphics   Graphics     `xml:"graphics"`
	Video      Video        `xml:"video"`
	Memballoon Memballoon   `xml:"memballoon"`
}

type Disk struct {
	Type    string  `xml:"type,attr"`
	Device  string  `xml:"device,attr"`
	Driver  Driver  `xml:"driver"`
	Source  Source  `xml:"source"`
	Target  Target  `xml:"target"`
	Address Address `xml:"address"`
}

type Driver struct {
	Name  string `xml:"name,attr"`
	Type  string `xml:"type,attr"`
	Cache string `xml:"cache,attr"`
}

type Source struct {
	File string `xml:"file,attr"`
}

type Target struct {
	Dev string `xml:"dev,attr"`
	Bus string `xml:"bus,attr"`
}

type Address struct {
	Type     string `xml:"type,attr"`
	Domain   string `xml:"domain,attr"`
	Bus      string `xml:"bus,attr"`
	Slot     string `xml:"slot,attr"`
	Function string `xml:"function,attr"`
}

type Controller struct {
	Type    string  `xml:"type,attr"`
	Index   int     `xml:"index,attr"`
	Model   string  `xml:"model,attr"`
	Master  Master  `xml:"master,omitempty"`
	Address Address `xml:"address"`
}

type Master struct {
	StartPort int `xml:"startport,attr"`
}

type Interface struct {
	Type    string  `xml:"type,attr"`
	MAC     MAC     `xml:"mac"`
	Source  Source  `xml:"source"`
	Model   Model   `xml:"model"`
	Address Address `xml:"address"`
}

type MAC struct {
	Address string `xml:"address,attr"`
}

type Model struct {
	Type string `xml:"type,attr"`
}

type Serial struct {
	Type   string `xml:"type,attr"`
	Target Target `xml:"target"`
}

type Console struct {
	Type   string `xml:"type,attr"`
	Target Target `xml:"target"`
}

type Channel struct {
	Type    string  `xml:"type,attr"`
	Source  Source  `xml:"source"`
	Target  Target  `xml:"target"`
	Address Address `xml:"address"`
}

type Input struct {
	Type    string  `xml:"type,attr"`
	Bus     string  `xml:"bus,attr"`
	Address Address `xml:"address"`
}

type Graphics struct {
	Type     string `xml:"type,attr"`
	Port     int    `xml:"port,attr"`
	Autoport string `xml:"autoport,attr"`
	Listen   Listen `xml:"listen"`
}

type Listen struct {
	Type    string `xml:"type,attr"`
	Address string `xml:"address,attr"`
}

type Video struct {
	Model   Model   `xml:"model"`
	Address Address `xml:"address"`
}

type Memballoon struct {
	Model   string  `xml:"model,attr"`
	Address Address `xml:"address"`
}

func TestXML(t *testing.T) {

	xmlFile, err := ioutil.ReadFile("domain.xml")
	if err != nil {
		log.Fatal(err)
	}

	var domain Domain
	err = xml.Unmarshal(xmlFile, &domain)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", domain)
}
