package main

type stater interface {
	doClock(contexter, int)
	doUse(contexter)
	doAlarm(contexter)
	doPhone(contexter)
}
