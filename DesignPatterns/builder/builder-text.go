package main

import (
	"fmt"
	"strings"
)

type markdownBuilder struct {
	buffer strings.Builder
	title  string
}

func (tb *markdownBuilder) makeTitle(title string) {
	tb.title = title
	tb.buffer.WriteString(
		fmt.Sprintf("# %s\n", title),
	)
	tb.buffer.WriteString("======\n")
}

func (tb *markdownBuilder) makeString(s string) {
	tb.buffer.WriteString(s)
	tb.buffer.WriteString("\n")
}

func (tb *markdownBuilder) makeItems(strs []string) {
	for _, s := range strs {
		tb.buffer.WriteString("  - ")
		tb.buffer.WriteString(s)
		tb.buffer.WriteString("\n")
	}
}

func (tb *markdownBuilder) output() {
	fmt.Println(tb.buffer.String())
}
