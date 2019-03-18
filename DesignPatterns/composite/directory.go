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

func (d *directory) print(path string) {
	path = fmt.Sprintf("%s/%s", path, d.name)
	fmt.Printf("%s (%d)\n", path, d.getSize())
	for _, e := range d.contents {
		e.print(path)
	}
}

func (d *directory) add(n entry) {
	d.contents = append(d.contents, n)
}
