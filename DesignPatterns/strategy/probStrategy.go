package main

import "math/rand"

type probStrategy struct {
	pre, cur int // hand index
	history  [][]int
}

func newProbStrategy() *probStrategy {
	return &probStrategy{
		history: [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
	}
}

func (ps *probStrategy) nextHand() *hand {
	bet := rand.Intn(ps.getSum(ps.cur))
	var i int
	switch {
	case bet < ps.history[ps.cur][0]:
		i = 0
	case bet < ps.history[ps.cur][0]+ps.history[ps.cur][1]:
		i = 1
	default:
		i = 2
	}
	ps.pre, ps.cur = ps.cur, i
	return getHand(i)
}

func (ps *probStrategy) study(win bool) {
	if win {
		ps.history[ps.pre][ps.cur]++
		return
	}
	ps.history[ps.pre][(ps.cur+1)%3]++
	ps.history[ps.pre][(ps.cur+2)%3]++
}

func (ps *probStrategy) getSum(g int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += ps.history[g][i]
	}
	return sum
}
