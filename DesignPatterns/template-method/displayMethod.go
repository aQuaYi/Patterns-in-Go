package main

type displayMethod struct {
	open, print, close func()
}

// display 方法规定了整个显示流程
func (r displayMethod) display() {
	r.open()
	for i := 0; i < 3; i++ {
		r.print()
	}
	r.close()
}
