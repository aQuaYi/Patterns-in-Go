package main

type display struct {
	displayImp
}

func newDisplay(di displayImp) *display {
	return &display{
		displayImp: di,
	}
}

func (d *display) play() {
	d.open()
	d.print()
	d.close()
}
