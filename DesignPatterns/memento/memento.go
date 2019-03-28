package main

import (
	"fmt"
	"strings"
)

type memento struct {
	money  int
	fruits []string
}

func createMemento(g *gamer) *memento {
	fruits := make([]string, 0, len(g.fruits))
	for _, f := range g.fruits {
		if strings.HasPrefix(f, "好吃的") { // 仅记录 “好吃的” 水果
			fruits = append(fruits, f)
		}
	}
	fmt.Println("in createMemento", fruits)
	return &memento{
		money:  g.money,
		fruits: fruits,
	}
}

func (m *memento) getMoney() int {
	return m.money
}
