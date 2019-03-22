package main

import "fmt"

// listVisitor implements interface:
// - visitor
// - visitorMethod

type listVisitor struct {
	currentDir string
}

func newListVisitor(path string) *listVisitor {
	return &listVisitor{
		currentDir: path,
	}
}

func (l *listVisitor) visit(inf interface{}) {
	switchMethod(l, inf)
}

func (l *listVisitor) visitFile(f file) {
	fmt.Printf("%s/%s", l.currentDir, f)
}
