package main

import "fmt"

type file struct {
	name string
	size int
}

func newFile(name string, size int) *file {
	return &file{
		name: name,
		size: size,
	}
}

func (f *file) getName() string {
	return f.name
}

func (f *file) getSize() int {
	return f.size
}

func (f *file) print(path string) {
	fmt.Printf("%s/%s (%d)\n", path, f.name, f.size)
}

func (f *file) add(n entry) {
	panic("can NOT add into file.")
}
