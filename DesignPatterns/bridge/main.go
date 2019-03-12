package main

func main() {
	d1 := newDisplay(newStringDisplayImp("Hello, China."))
	d1.play()

	d2 := newCountDisplay(newStringDisplayImp("Hello, World."))
	d2.play()

	d3 := newCountDisplay(newStringDisplayImp("Hello, Universe."))
	d3.play()
	d3.multiDisplay(3)

	d4 := newIncreaseDisplay(newCharDisplayImp('<', '*', '>'), 1)
	d4.increaseDisplay(4)

	d5 := newIncreaseDisplay(newCharDisplayImp('|', '#', '-'), 2)
	d5.increaseDisplay(6)

	d4.play()

	d5.play()

}
