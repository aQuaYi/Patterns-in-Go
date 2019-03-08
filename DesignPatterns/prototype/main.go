package main

const (
	strongMessage = "strong message"
	warningBox    = "warning box"
	slashBox      = "slash box"
)

func main() {
	manager := newManager()

	strongMsg := newUnderlinePen('~')
	manager.register(strongMessage, strongMsg)
	wBox := newMessageBox('*')
	manager.register(warningBox, wBox)
	sBox := newMessageBox('/')
	manager.register(slashBox, sBox)

	// manager 只需要知道对象的代号，就可以复制出同样的对象。
	// 这样就与对象的具体类型解耦了
	p1 := manager.create(strongMessage)
	p1.use("the first line")

	p2 := manager.create(warningBox)
	p2.use("the second line")

	p3 := manager.create(slashBox)
	p3.use("hello, world.")
}
