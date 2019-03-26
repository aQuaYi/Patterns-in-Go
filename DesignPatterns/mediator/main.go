package main

// 由于 oven 功耗很高，开启后，必须关闭 computer 和 refrigerator
// oven 关闭后，自动开启 refrigerator
// fan 保持与 computer 状态一致

const (
	on  = true
	off = false
)

func main() {
	remote := newRemote()
	computer, oven, fan := remote.config()

	fan.turn(on)
	fan.turn(off)
	fan.turn(on)

	computer.turn(on)
	computer.turn(off)
	computer.turn(on)

	oven.turn(on)
	oven.turn(off)
	oven.turn(on)

}
