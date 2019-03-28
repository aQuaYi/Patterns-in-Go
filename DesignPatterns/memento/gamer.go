package main

import (
	"fmt"
	"math/rand"
)

type gamer struct {
	money      int
	fruits     []string
	fruitsName []string
}

func newGamer(money int) *gamer {
	return &gamer{
		money:      money,
		fruitsName: []string{"苹果", "葡萄", "香蕉", "橘子"},
	}
}

func (g *gamer) getMoney() int {
	return g.money
}

func (g *gamer) randFruit() string {
	prefix := ""
	if rand.Int()%2 == 0 {
		prefix = "好吃的"
	}
	return prefix + g.fruitsName[rand.Intn(len(g.fruitsName))]
}

func (g *gamer) createMemento() *memento {
	return createMemento(g)
}

func (g *gamer) restoreMemento(m *memento) {
	g.money = m.money
	g.fruits = m.fruits
}

func (g *gamer) String() string {
	return fmt.Sprintf("[money=%d, fruits=%v]", g.money, g.fruits)
}

func (g *gamer) bet() {
	switch rand.Intn(6) + 1 {
	case 1:
		g.money += 100
		fmt.Println("金钱 +100")
	case 2:
		g.money /= 2
		fmt.Println("金钱减半")
	case 6:
		f := g.randFruit()
		fmt.Printf("获得了水果(%s)\n", f)
		g.fruits = append(g.fruits, f)
	default:
		fmt.Println("无事发生")
	}
}
