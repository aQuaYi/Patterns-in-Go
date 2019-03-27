package main

import "fmt"

func main() {
	pub := newNumberGenerator()

	d := newDigit()
	pub.add(d)
	g := newGraph()
	pub.add(g)

	pub.execute(3)

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-")

	pub.remove(d)
	pub.execute(3)
}
