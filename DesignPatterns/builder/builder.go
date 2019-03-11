package main

type builder interface {
	makeTitle(string)
	makeString(string)
	makeItems([]string)
	output()
}
