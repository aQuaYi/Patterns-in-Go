package main

import (
	"fmt"
	"strings"
)

type markdownBuilder struct {
	buffer          strings.Builder
	filename, title string
}

func newMarkdownBuilder() *markdownBuilder {
	return &markdownBuilder{}
}

func (m *markdownBuilder) begin(filename string) {
	m.filename = filename + ".md"
}

func (m *markdownBuilder) end() {
}

func (m *markdownBuilder) makeTitle(title string) {
	m.title = title
	m.buffer.WriteString(
		fmt.Sprintf("# %s\n", title),
	)
	m.buffer.WriteString("======\n")
}

func (m *markdownBuilder) makeString(s string) {
	m.buffer.WriteString(s)
	m.buffer.WriteString("\n")
}

func (m *markdownBuilder) makeItems(strs []string) {
	for _, s := range strs {
		m.buffer.WriteString("  - ")
		m.buffer.WriteString(s)
		m.buffer.WriteString("\n")
	}
}

func (m *markdownBuilder) output() {
	fmt.Printf("The file is %s\n", m.filename)
	fmt.Println(m.buffer.String())
}
