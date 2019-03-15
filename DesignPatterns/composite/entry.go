package main

type entry interface {
	getName() string
	getSize() int
	print()
	add(entry)
}
