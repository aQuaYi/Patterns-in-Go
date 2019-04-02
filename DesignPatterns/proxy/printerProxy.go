package main

import "fmt"

type printerProxy struct {
	name string
	real printable
}

func newPrinterProxy(name string) *printerProxy {
	fmt.Printf("new printer proxy: %s\n", name)
	return &printerProxy{
		name: name,
	}
}

func (p *printerProxy) getName() string {
	return p.name
}

func (p *printerProxy) setName(name string) {
	fmt.Printf("printer proxy change name to %s\n", name)
	p.name = name
	if p.real != nil {
		p.real.setName(name)
	}
}

func (p *printerProxy) print(s string) {
	if p.real == nil {
		// 把真实 printer 的生成工作，推迟到调用 print() 时
		p.real = newPrinter(p.name)
	}
	p.real.print(s)
}
