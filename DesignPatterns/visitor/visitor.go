package main

type visitor interface {
	visit(interface{})
}

type visitorMethod interface {
	visitFile(file)
	visitDirectory(directory)
}

func switchMethod(v visitor, inf interface{}) {
	vm, ok := v.(visitorMethod)
	if !ok {
		panic("v do not implement visitorMethod interface")
	}
	switch inf.(type) {
	case file:
		f := inf.(file)
		vm.visitFile(f)
	case directory:
		d := inf.(directory)
		vm.visitDirectory(d)
	default:
		panic("inf is not an entry object")
	}
}
