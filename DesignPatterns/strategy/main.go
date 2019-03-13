package main

import "fmt"

func main() {
	winnie := newPlayer("winnie", newWinningStrategy())
	pope := newPlayer("pope", newProbStrategy())

	for i := 0; i < 1000000; i++ {
		wHand := winnie.showHand()
		pHand := pope.showHand()
		winnie.check(pHand)
		pope.check(wHand)
	}

	fmt.Println("Result:")
	fmt.Println(winnie)
	fmt.Println(pope)
}
