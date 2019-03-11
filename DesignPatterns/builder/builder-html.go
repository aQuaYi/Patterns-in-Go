package main

import (
	"fmt"
	"strings"
)

type htmlBuilder struct {
	buffer          strings.Builder
	filename, title string
}

func newHTMLBuilder() *htmlBuilder {
	return &htmlBuilder{}
}

func (h *htmlBuilder) begin(filename string) {
	h.filename = filename + ".html"
	h.buffer.WriteString(
		"<html>\n<head>\n<title>" +
			filename +
			"</title>\n</head>\n<body>\n",
	)
}

func (h *htmlBuilder) end() {
	h.buffer.WriteString(
		"</body>\n</html>\n",
	)
}

func (h *htmlBuilder) makeTitle(title string) {
	h.title = title
	h.buffer.WriteString(
		fmt.Sprintf("<h1>%s</h1>\n", title),
	)
}

func (h *htmlBuilder) makeString(str string) {
	h.buffer.WriteString(
		fmt.Sprintf("<p>%s</p>\n", str),
	)
}

func (h *htmlBuilder) makeItems(strs []string) {
	h.buffer.WriteString("<ul>\n")
	for _, s := range strs {
		h.buffer.WriteString(
			"<li>" + s + "</li>\n",
		)
	}
	h.buffer.WriteString("</ul>\n")
}

func (h *htmlBuilder) output() {
	fmt.Printf("The file is %s\n", h.filename)
	fmt.Println(h.buffer.String())
}
