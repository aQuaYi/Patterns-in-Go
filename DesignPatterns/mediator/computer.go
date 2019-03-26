package main

import "fmt"

// implements colleaguer
type computer struct {
	mediator
	isWorking bool
}

func newComputer() *computer {
	return &computer{}
}

func (c *computer) setMediator(m mediator) {
	c.mediator = m
}

func (c *computer) setColleagueEnable(enabled bool) {
	switch {
	case c.isWorking && enabled:
		fmt.Println("computer")
	case c.isWorking && !enabled:
	case !c.isWorking && enabled:
	default: // !c.isWorking && !enabled
	}
}
