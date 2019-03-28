package main

import "fmt"

type safe struct {
	stater
}

func newSafe() *safe {
	return &safe{
		stater: newNightState(), // 默认黑夜
	}
}

func (s *safe) setClock(hour int) {
	s.doClock(s, hour)
	fmt.Printf("现在时间是 %02d:00, %s\n", hour, s.stater)
}

func (s *safe) changeState(state stater) {
	fmt.Printf("\t%s -> %s\n", s.stater, state)
	s.stater = state
}

func (s *safe) callSecurityCenter(msg string) {
	fmt.Printf("Call  : %s\n", msg)
}

func (s *safe) recordLog(msg string) {
	fmt.Printf("Record: %s\n", msg)
}
