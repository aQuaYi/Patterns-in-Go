package main

type stringDisplay struct {
	display
	string
}

func newStringDisplay(s string) *stringDisplay {
	return &stringDisplay{
		display: display{},
		string:  s,
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
