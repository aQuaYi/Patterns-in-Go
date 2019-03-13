package main

import "fmt"

type player struct {
	name   string
	myHand *hand
	strategy
	win, total int // count
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
	win := p.myHand.isWin(h)
	p.strategy.study(win)
	if win {
		p.win++
	}
	p.total++
}

func (p *player) String() string {
	format := "%8s: %d win in %d games"
	return fmt.Sprintf(format, p.name, p.win, p.total)
}
