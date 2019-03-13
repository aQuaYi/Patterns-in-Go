package main

import "fmt"

type player struct {
	name   string
	myHand *hand
	strategy
	win, lose, total int // count
}

func newPlayer(name string, strategy strategy) *player {
	return &player{
		name:     name,
		strategy: strategy,
	}
}

func (p *player) showHand() *hand {
	p.myHand = p.strategy.nextHand()
	return p.myHand
}

func (p *player) check(h *hand) {
	result := p.myHand.fight(h)
	p.strategy.study(result == 1)
	switch result {
	case 1:
		p.win++
	case -1:
		p.lose++
	}
	p.total++
}

func (p *player) String() string {
	format := "%8s: %d win and %d lose in %d games"
	return fmt.Sprintf(format, p.name, p.win, p.lose, p.total)
}
