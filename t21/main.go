package main

import "fmt"

// TODO| Реализовать паттерн «адаптер» на любом примере.

type USBTypeC interface {
	InsertTypeC()
}

type Phone struct{}

func (p *Phone) InsertTypeC() {
	fmt.Println("Inserting usb Type-C to phone")
}

type Computer struct{}

func (c *Computer) InsertUSB() {
	fmt.Println("Inserting usb to computer")
}

type USBToTypeC struct {
	comp Computer
}

func (a *USBToTypeC) InsertTypeC() {
	fmt.Println("Inserting usb Type-C to computer")
}

func main() {
	computer := Computer{}
	var adapter USBTypeC = &USBToTypeC{computer}
	adapter.InsertTypeC()
}
