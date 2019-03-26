package main

import "fmt"

// implements colleaguer
type fan struct {
	mediator
	isWorking bool
	tab       string
}

func newFan() *fan {
	return &fan{}
}

func (f *fan) setMediator(m mediator) {
	f.mediator = m
}

func (f *fan) turn(enabled bool) {
	if enabled && !f.isWorking {
		fmt.Println(f.tab + "fan is on...")
	} else if !enabled && f.isWorking {
		fmt.Println(f.tab + "fan is off...")
	}
	f.tab = ""
	f.isWorking = enabled
}

func (f *fan) setStatus(enabled bool) {
	f.tab = "\t"
	f.turn(enabled)
}
