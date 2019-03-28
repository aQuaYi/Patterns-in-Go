package main

import "sync"

var (
	dayOnce     sync.Once
	dayInstance *dayState
)

type dayState struct {
}

func newDayState() *dayState {
	dayOnce.Do(func() {
		dayInstance = &dayState{}
	})
	return dayInstance
}

func (d *dayState) doClock(c contexter, hour int) {
	if hour < 9 || 17 <= hour {
		c.changeState(newNightState())
	}
}

func (d *dayState) doUse(c contexter) {
	c.recordLog("使用金库（白天）")
}

func (d *dayState) doAlarm(c contexter) {
	c.callSecurityCenter("按下警铃（白天）")
}

func (d *dayState) doPhone(c contexter) {
	c.callSecurityCenter("正常通话（白天）")
}

func (d *dayState) String() string {
	return "[白天]"
}
