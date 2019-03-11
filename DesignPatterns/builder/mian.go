package main

func main() {
	mdb := newMarkdownBuilder()
	director := newDirector(mdb)
	director.construct()

	hb := newHTMLBuilder()
	director = newDirector(hb)
	director.construct()
}
