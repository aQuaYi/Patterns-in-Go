package main

import "fmt"

type book struct {
	name string
}

func newBook(name string) book {
	return book{
		name: name,
	}
}

func (b book) String() string {
	return fmt.Sprintf("book: %s", b.name)
}
