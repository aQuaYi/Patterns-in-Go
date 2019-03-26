package main

import "fmt"

// implements colleaguer
type refrigerator struct {
	mediator
	isWorking bool
}

func newRefrigerator() *refrigerator {
	return &refrigerator{}
}

func (r *refrigerator) setMediator(m mediator) {
	r.mediator = m
}

func (r *refrigerator) setStatus(enabled bool) {
	r.isWorking = enabled
	if enabled {
		fmt.Println("\trefrigerator is on...")
	} else {
		fmt.Println("\trefrigerator is off...")
	}
}
