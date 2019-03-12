package main

func main() {
	d1 := newDisplay(newStringDisplayImp("Hello, China."))
	d1.play()

	d2 := newDisplay(newStringDisplayImp("Hello, World."))
	d2.play()

	d3 := newCountDisplay(newStringDisplayImp("Hello, Universe."))
	d3.play()
	d3.multiDisplay(3)

}
