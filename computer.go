package main

// 抽象层
type Computer interface {
	Print()
	SetPrinter(Printer)
}
