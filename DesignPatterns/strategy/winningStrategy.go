package main

import (
	"math/rand"
)

type winningStrategy struct {
	won     bool
	preHand *hand
}

func newWinningStrategy() *winningStrategy {
	return &winningStrategy{
		won:     false,
		preHand: nil,
	}
}

func (ws *winningStrategy) nextHand() *hand {
	if !ws.won {
		ws.preHand = getHand(rand.Intn(3))
	}
	return ws.preHand
}

func (ws *winningStrategy) study(win bool) {
	ws.won = win
}
