package main

import "fmt"

type trouble struct {
	num int
}

func newTrouble(n int) *trouble {
	return &trouble{
		num: n,
	}
}

func (t *trouble) getNumber() int {
	return t.num
}

func (t *trouble) String() string {
	return fmt.Sprintf("[trouble:%d]", t.num)
}
