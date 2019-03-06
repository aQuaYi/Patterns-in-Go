package main

import (
	"fmt"
	"strings"
)

type str struct {
	str string
	displayMethod
}

func newStr(s string) *str {
	sp := &str{
		str: s,
	}

	line := strings.Repeat("-", len(s)+2)

	// str.displayMethod.open() 在这里被个性化的定义
	// 后面的 print() 和 close() 方法同理
	sp.open = func() {
		fmt.Printf("+%s+\n", line)
	}

	sp.print = func() {
		fmt.Printf("| %s |\n", sp.str)
	}

	sp.close = func() {
		fmt.Printf("+%s+", line)
	}

	return sp
}
