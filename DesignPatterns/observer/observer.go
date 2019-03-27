package main

import (
	"fmt"
	"strings"
)

type observer interface {
	update(publisher)
}

type digit struct {
}

func newDigit() *digit {
	return &digit{}
}

func (d *digit) update(p publisher) {
	fmt.Printf("digit:%d\n", p.getNumber())
}

type graph struct {
}

func newGraph() *graph {
	return &graph{}
}

func (g *graph) update(p publisher) {
	stars := strings.Repeat("*", p.getNumber())
	fmt.Printf("graph:%s\n", stars)
}
