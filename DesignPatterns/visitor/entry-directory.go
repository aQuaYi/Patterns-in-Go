package main

import "fmt"

type directory struct {
	name     string
	size     int
	contents []entry
}

func newDirectory(name string) *directory {
	return &directory{
		name: name,
	}
}

func (d *directory) String() string {
	return fmt.Sprintf("%s (%d)", d.getName(), d.getSize())
}

func (d *directory) getName() string {
	return d.name
}

func (d *directory) getSize() int {
	size := 0
	for _, c := range d.contents {
		size += c.getSize()
	}
	return size
}

func (d *directory) add(n entry) {
	d.contents = append(d.contents, n)
}

func (d *directory) accept(v visitor) {
	v.visit(d)
}
