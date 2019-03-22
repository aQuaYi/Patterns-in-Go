package main

type entry interface {
	getName() string
	getSize() int
	add(entry)
	accept(visitor)
}
