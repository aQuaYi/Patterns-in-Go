package main

import "fmt"

// listVisitor implements interface:
// - visitor
// - visitorMethod

type listVisitor struct {
	currentDir string
}

func newListVisitor() *listVisitor {
	return &listVisitor{}
}

func (l *listVisitor) visit(inf interface{}) {
	switchMethod(l, inf)
}

func (l *listVisitor) visitFile(f *file) {
	fmt.Printf("%s/%s\n", l.currentDir, f)
}

func (l *listVisitor) visitDirectory(d *directory) {
	fmt.Printf("%s/%s\n", l.currentDir, d)
	oldDir := l.currentDir
	l.currentDir += "/" + d.getName()
	for _, c := range d.contents {
		c.accept(l)
	}
	l.currentDir = oldDir
}
