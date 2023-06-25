package main

import "fmt"

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for Mac")
}

func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}

type Windows struct {
	printer Printer
}

func (m *Windows) Print() {
	fmt.Print("Print request for Windows")
}

func (w *Windows) SetPrinter(printer Printer) {
	w.printer = printer
}
