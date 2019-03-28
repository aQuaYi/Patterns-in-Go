package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	safe := newSafe()
	for i := 0; i < 48; i++ {
		hour := i % 24
		if hour == 0 {
			fmt.Println("-=-=-=-=-=-=-=-")
		}

		safe.setClock(hour)

		state := safe.stater
		switch rand.Intn(5) {
		case 0:
			state.doAlarm(safe)
		case 1:
			state.doUse(safe)
		case 2:
			state.doPhone(safe)
		}
	}
}
