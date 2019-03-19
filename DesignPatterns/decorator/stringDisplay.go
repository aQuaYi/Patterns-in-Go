package main

// stringDisplay implements display interface
type stringDisplay struct {
	string
}

func newStringDisplay(s string) *stringDisplay {
	return &stringDisplay{
		string: s,
	}
}

func (sd *stringDisplay) getColumns() int {
	return len(sd.string)
}

func (sd *stringDisplay) getRows() int {
	return 1
}

func (sd *stringDisplay) getRowText(row int) string {
	return sd.string
}

func (sd *stringDisplay) show() {
	showFunc(sd)
}
