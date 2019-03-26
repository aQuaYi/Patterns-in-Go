package main

import "fmt"

// implements colleaguer
type oven struct {
	mediator
	isWorking bool
}

func newOven() *oven {
	return &oven{}
}

func (o *oven) setMediator(m mediator) {
	o.mediator = m
}

func (o *oven) setStatus(enabled bool) {
	o.isWorking = enabled
	if enabled {
		fmt.Println("\toven is on...")
	} else {
		fmt.Println("\toven is off...")
	}
}

func (o *oven) turn(enabled bool) {
	if enabled {
		fmt.Println("oven is on...")
	} else {
		fmt.Println("oven is off...")
	}
	o.mediator.turnedOven(enabled)
}
