package main

type gesture int

const (
	// ROCK 代表石头
	ROCK gesture = iota
	// SCISSORS 代表剪刀
	SCISSORS
	// PAPER 代表布
	PAPER
)

var hands = []*hand{
	newHand(ROCK),
	newHand(SCISSORS),
	newHand(PAPER),
}

func getHand(gi int) *hand {
	return hands[gi]
}

type hand struct {
	gesture
}

func newHand(g gesture) *hand {
	return &hand{
		gesture: g,
	}
}

// return 1 if win
func (h *hand) isWin(opponent *hand) bool {
	if (h.gesture+1)%3 == opponent.gesture {
		return true
	}
	return false
}
