package main

import "fmt"

// implements colleaguer
type computer struct {
	mediator
	isWorking bool
	tab       string
}

func newComputer() *computer {
	return &computer{}
}

func (c *computer) setMediator(m mediator) {
	c.mediator = m
}

func (c *computer) setStatus(enabled bool) {
	c.tab = "\t"
	c.turn(enabled)
}

func (c *computer) turn(enabled bool) {
	if enabled && !c.isWorking {
		fmt.Println(c.tab + "computer is on...")
	} else if !enabled && c.isWorking {
		fmt.Println(c.tab + "computer is off...")
	}
	c.tab = ""
	c.isWorking = enabled
	c.mediator.turnedComputer(enabled)
}
