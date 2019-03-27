package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type publisher interface {
	getNumber() int
}

type randNumberGenerator struct {
	observers map[observer]struct{}
	number    int
}

func newNumberGenerator() *randNumberGenerator {
	return &randNumberGenerator{
		observers: make(map[observer]struct{}, 128),
	}
}

func (g *randNumberGenerator) add(o observer) {
	g.observers[o] = struct{}{}
}

func (g *randNumberGenerator) remove(o observer) {
	delete(g.observers, o)
}

func (g *randNumberGenerator) getNumber() int {
	return g.number
}

func (g *randNumberGenerator) notify() {
	for o := range g.observers {
		o.update(g)
	}
}

func (g *randNumberGenerator) execute(n int) {
	for i := 0; i < n; i++ {
		g.number = rand.Intn(50)
		g.notify()
	}
}
