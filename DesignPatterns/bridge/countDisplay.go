package main

type countDisplay struct {
	*display
}

func newCountDisplay(d displayImp) *countDisplay {
	return &countDisplay{
		display: newDisplay(d),
	}
}

func (cd *countDisplay) multiDisplay(n int) {
	cd.open()
	for i := 0; i < n; i++ {
		cd.print()
	}
	cd.close()
}
