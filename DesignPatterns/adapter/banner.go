package main

import "fmt"

type shower interface {
	show() // show original string
	showWithParen()
	showWithAster()
}

// banner implements shower interface
type banner struct {
	string
}

func newBanner(content string) *banner {
	return &banner{
		string: content,
	}
}

func (b *banner) show() {
	fmt.Print(b.string)
}

func (b *banner) showWithParen() {
	fmt.Printf("(%s)", b.string)
}

func (b *banner) showWithAster() {
	fmt.Printf("*%s*", b.string)
}
