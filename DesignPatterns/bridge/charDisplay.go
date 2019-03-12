package main

import "fmt"

type charDisplayImp struct {
	head, body, tail byte
}

func newCharDisplayImp(head, body, tail byte) *charDisplayImp {
	return &charDisplayImp{
		head: head,
		body: body,
		tail: tail,
	}
}

func (cdi *charDisplayImp) open() {
	fmt.Printf("%c", cdi.head)
}

func (cdi *charDisplayImp) print() {
	fmt.Printf("%c", cdi.body)
}

func (cdi *charDisplayImp) close() {
	fmt.Printf("%c\n", cdi.tail)
}
