package main

import "time"

func main() {
	p := newPrinterProxy("Alice")

	p.setName("Bob")
	p.print(time.Now().String())

	p.setName("Cindy")
	p.print(time.Now().String())
}
