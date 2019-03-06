package main

type printShower struct {
	shower
}

func newPrintShower(s shower) *printShower {
	return &printShower{
		shower: s,
	}
}

func (pb *printShower) print() {
	// NOTICE:
	// printShower 作为适配器的作用在于委托
	// 把 print() 的具体工作委托给 shower.show() 完成
	// printWeak() 与 printStrong() 同理
	pb.shower.show()
}

func (pb *printShower) printWeak() {
	pb.shower.showWithParen()
}

func (pb *printShower) printStrong() {
	pb.shower.showWithAster()
}
