package main

import "fmt"

type Epson struct{}

func (e Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type HP struct{}

func (h HP) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
