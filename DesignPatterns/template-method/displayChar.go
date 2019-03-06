package main

import "fmt"

type char struct {
	char byte
	displayMethod
}

func newChar(c byte) *char {
	cp := &char{
		char: c,
	}

	// char.displayMethod.open() 在这里被个性化的定义
	// 后面的 print() 和 close() 方法同理
	cp.open = func() {
		fmt.Print("<<")
	}

	cp.print = func() {
		fmt.Printf("%c", cp.char)
	}

	cp.close = func() {
		fmt.Print(">>")
	}

	return cp
}
