package main

type strategy interface {
	nextHand() *hand
	study(win bool)
}
