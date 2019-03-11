package main

type builder interface {
	begin(string)
	makeTitle(string)
	makeString(string)
	makeItems([]string)
	end()
	output()
}
