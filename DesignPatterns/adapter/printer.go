package main

import "fmt"

type printer interface {
	print()
	printWeak()
	printStrong()
}

func print(p printer) {
	fmt.Print("original: ")
	p.print()
	fmt.Println()
	fmt.Print("weak    : ")
	p.printWeak()
	fmt.Println()
	fmt.Print("strong  : ")
	p.printStrong()
	fmt.Println()
}
