package main

type safe struct {
	state stater
}

func newSafe() *safe {
	return &safe{
		state: newDayState(), // 默认白天
	}
}

func (s *safe) setClock(hour int) {

}
