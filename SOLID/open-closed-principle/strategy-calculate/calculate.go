package main

import "fmt"

type calculate struct {
	a, b int
	executer
}

type executer interface {
	execute(int, int) int
}

func newCalculate(a, b int, e executer) *calculate {
	return &calculate{
		a:        a,
		b:        b,
		executer: e,
	}
}

func (c *calculate) do() int {
	return c.execute(c.a, c.b)
}

type add struct {
}

func (a add) execute(i, j int) int {
	return i + j
}

type multiply struct {
}

func (m multiply) execute(a, b int) int {
	return a * b
}

func main() {
	c1 := newCalculate(2, 3, add{})
	fmt.Println("calculate result is ", c1.do())

	c2 := newCalculate(2, 3, multiply{})
	fmt.Println("calculate result is ", c2.do())
}
