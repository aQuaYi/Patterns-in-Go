package main

type printable interface {
	getName() string
	setName(string)
	print(string)
}
