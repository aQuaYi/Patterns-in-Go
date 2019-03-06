package main

import "fmt"

func main() {

	/**
	 * char 和 str 使用同样的 display 流程
	 * 由 displayMethod.display() 来定义
	 * 但是，displayMethod.display() 中用到的 open(), print(), close()
	 * 却是由 char 和 str 各自定义的。
	 */

	char := newChar('H')
	char.display()

	fmt.Println()

	str := newStr("hello world")
	str.display()

	fmt.Println()
}
