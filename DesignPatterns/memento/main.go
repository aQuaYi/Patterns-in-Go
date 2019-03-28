package main

import (
	"fmt"
	"time"
)

func main() {
	gamer := newGamer(100)
	m := gamer.createMemento()
	for i := 0; i < 100; i++ {
		fmt.Printf("%02d == %s ==\n", i, gamer)

		gamer.bet()

		if gamer.getMoney() > m.getMoney() {
			fmt.Println("（金钱增加，游戏存档...）")
			m = gamer.createMemento()
		} else if gamer.getMoney() < m.getMoney()/4 {
			fmt.Println("（金钱减少3/4多，恢复存档...）")
			gamer.restoreMemento(m)
		}

		time.Sleep(time.Second)
	}
}
