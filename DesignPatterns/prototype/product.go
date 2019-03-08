package main

import (
	"fmt"
	"strings"
)

type product interface {
	use(string)
	createClone() product
}

// messageBox implements product interface
type messageBox struct {
	decochar byte
}

func newMessageBox(char byte) *messageBox {
	return &messageBox{
		decochar: char,
	}
}

func (mb *messageBox) use(s string) {
	line := strings.Repeat(string(mb.decochar), len(s)+4)
	fmt.Println(line)
	fmt.Printf("%c %s %c\n", mb.decochar, s, mb.decochar)
	fmt.Println(line)
}

func (mb *messageBox) createClone() product {
	clone := *mb // shallow copy
	return &clone
}

// underlinePen implements product interface
type underlinePen struct {
	ulchar byte
}

func newUnderlinePen(char byte) *underlinePen {
	return &underlinePen{
		ulchar: char,
	}
}

func (ul *underlinePen) use(s string) {
	fmt.Printf("\"%s\"\n", s)
	line := strings.Repeat(string(ul.ulchar), len(s))
	fmt.Printf(" %s \n", line)
}

func (ul *underlinePen) createClone() product {
	clone := *ul // shallow copy
	return &clone
}
