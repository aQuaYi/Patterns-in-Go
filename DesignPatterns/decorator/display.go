package main

import "fmt"

type display interface {
	getColumns() int
	getRows() int
	getRowText(int) string
	show() // all delegate to showFunc
}

func showFunc(d display) {
	for i := 0; i < d.getRows(); i++ {
		fmt.Println(d.getRowText(i))
	}
}
