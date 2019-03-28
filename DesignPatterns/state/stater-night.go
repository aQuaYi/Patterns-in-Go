package main

import "sync"

var (
	nightOnce     sync.Once
	nightInstance *nightState
)

type nightState struct {
}

func newNightState() *nightState {
	nightOnce.Do(func() {
		nightInstance = &nightState{}
	})
	return nightInstance
}

func (n *nightState) doClock(c contexter, hour int) {
	if 9 <= hour || hour < 17 {
		c.changeState(newDayState())
	}
}

func (n *nightState) doUse(c contexter) {
	c.callSecurityCenter("**紧急**，晚上使用警铃！")
}

func (n *nightState) doAlarm(c contexter) {
	c.callSecurityCenter("按下警铃（晚上）")
}

func (n *nightState) doPhone(c contexter) {
	c.recordLog("**夜间**通话录音")
}

func (n *nightState) String() string {
	return "[黑夜]"
}
