package main

import "fmt"

type calculate struct {
	a, b int
}

func newCalculate(a, b int) *calculate {
	return &calculate{
		a: a,
		b: b,
	}
}

func (c *calculate) do() int {
	return c.a + c.b
}

func main() {
	c := newCalculate(2, 3)
	fmt.Println(c.do())
}
