package main

func main() {
	mdb := newMarkdownBuilder()
	director := newDirector(mdb)
	director.construct()

	hb := newHtmlBuilder()
	director = newDirector(hb)
	director.construct()
}
