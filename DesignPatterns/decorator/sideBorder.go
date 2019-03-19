package main

type sideBorder struct {
	display
	borderChar string
}

func newSideBorder(d display, char byte) *sideBorder {
	return &sideBorder{
		display:    d,
		borderChar: string(char),
	}
}

func (s *sideBorder) getColumns() int {
	return 1 + s.display.getColumns() + 1
}

func (s *sideBorder) getRows() int {
	return s.display.getRows()
}

func (s *sideBorder) getRowText(i int) string {
	return s.borderChar + s.display.getRowText(i) + s.borderChar
}

func (s *sideBorder) show() {
	showFunc(s)
}
