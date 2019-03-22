package main

func main() {
	root := newDirectory("root")
	bin := newDirectory("bin")
	tmp := newDirectory("tmp")
	usr := newDirectory("usr")

	bin.add(newFile("vi", 1000))
	bin.add(newFile("latex", 2000))

	root.add(bin)
	root.add(tmp)
	root.add(usr)

	root.accept(newListVisitor())
}
