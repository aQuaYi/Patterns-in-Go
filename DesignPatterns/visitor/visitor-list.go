package main

type listVisitor struct {
	currentDir string
}

// listVisitor implements visitor interface{}
func (l *listVisitor) visit(interface{}) {
}
