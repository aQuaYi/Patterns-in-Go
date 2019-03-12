package main

import "fmt"

type stringDisplayImp struct {
	str string
}

func newStringDisplayImp(s string) *stringDisplayImp {
	return &stringDisplayImp{
		str: s,
	}
}

func (s *stringDisplayImp) open() {
	s.printLine()
}

func (s *stringDisplayImp) print() {
	fmt.Printf("|%s|\n", s.str)
}

func (s *stringDisplayImp) close() {
	s.open()
}

func (s *stringDisplayImp) printLine() {
	fmt.Print("+")
	for i := 0; i < len(s.str); i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
