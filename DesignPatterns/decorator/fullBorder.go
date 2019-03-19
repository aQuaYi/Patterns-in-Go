package main

import (
	"fmt"
	"strings"
)

type fullBorder struct {
	display
}

func newFullBorder(d display) *fullBorder {
	return &fullBorder{
		display: d,
	}
}

func (f *fullBorder) getColumns() int {
	return 1 + f.display.getColumns() + 1
}

func (f *fullBorder) getRows() int {
	return 1 + f.display.getRows() + 1
}

func (f *fullBorder) getRowText(i int) string {
	begin, end := 0, f.display.getRows()+1
	line := fmt.Sprintf("+%s+",
		strings.Repeat("-", f.display.getColumns()))
	switch i {
	case begin, end:
		return line
	default:
		return "|" + f.display.getRowText(i-1) + "|"
	}
}

func (f *fullBorder) show() {
	showFunc(f)
}
