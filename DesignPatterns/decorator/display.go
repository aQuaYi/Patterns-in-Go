package main

import "fmt"

type display struct {
	getColumns func() int
	getRows    func() int
	getRowText func() string
}

func (d display) show() {
	for i := 0; i < d.getRows(); i++ {
		fmt.Println(d.getRowText())
	}
}
