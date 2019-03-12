package main

type display struct {
	displayImp
}

func newDisplay(d displayImp) *display {
	return &display{
		displayImp: d,
	}
}

func (d *display) play() {
	d.open()
	d.print()
	d.close()
}
