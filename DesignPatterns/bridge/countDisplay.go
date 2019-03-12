package main

type countDisplay struct {
	*display
}

func newCountDisplay(di displayImp) *countDisplay {
	return &countDisplay{
		display: newDisplay(di),
	}
}

func (cd *countDisplay) multiDisplay(n int) {
	cd.open()
	for i := 0; i < n; i++ {
		cd.print()
	}
	cd.close()
}
