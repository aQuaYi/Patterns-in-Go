package main

import (
	"fmt"
	"time"
)

type printer struct {
	name string
}

func newPrinter(name string) *printer {
	start()
	return &printer{
		name: name,
	}
}

func start() {
	fmt.Print("startint printer. ")
	s := "But the real printer is hard to start..................."
	for i := range s {
		time.Sleep(time.Millisecond * 250)
		fmt.Print(s[i : i+1])
	}
	fmt.Println()
}

func (p *printer) setName(name string) {
	fmt.Printf("printer's new name is %s\n", name)
	p.name = name
}

func (p *printer) getName() string {
	return p.name
}

func (p *printer) print(s string) {
	fmt.Printf("=== %s ===\n", p.name)
	fmt.Println(s)
}
