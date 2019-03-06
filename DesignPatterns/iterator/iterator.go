package main

import "fmt"

type iterator interface {
	hasNext() bool
	next() interface{}
}

func printAll(it iterator) {
	for it.hasNext() {
		fmt.Println(it.next())
	}
}
