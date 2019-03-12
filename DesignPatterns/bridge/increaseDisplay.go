package main

type increaseDisplay struct {
	*countDisplay
	step int
}

func newIncreaseDisplay(di displayImp, step int) *increaseDisplay {
	return &increaseDisplay{
		countDisplay: newCountDisplay(di),
		step:         step,
	}
}

func (id *increaseDisplay) increaseDisplay(level int) {
	count := 0
	for i := 0; i < level; i++ {
		id.multiDisplay(count)
		count += id.step
	}
}
