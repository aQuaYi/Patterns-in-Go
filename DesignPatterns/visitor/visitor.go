package main

type visitor interface {
	visit(interface{})
}

type visitorMethod interface {
	visitFile(file)
	visitDirectory(directory)
}

func switchMethod(v visitor, inf interface{}) {
	// TODO: 完成
	return
}
