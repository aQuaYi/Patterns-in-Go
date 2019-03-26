package main

import "fmt"

// implements colleaguer
type fan struct {
	mediator
	isWorking bool
}

func newFan() *fan {
	return &fan{}
}

func (f *fan) setMediator(m mediator) {
	f.mediator = m
}

func (f *fan) setStatus(enabled bool) {
	if enabled && !f.isWorking {
		fmt.Println("\tfan is on...")
	} else if !enabled && f.isWorking {
		fmt.Println("\tfan is off...")
	}
	f.isWorking = enabled
}
