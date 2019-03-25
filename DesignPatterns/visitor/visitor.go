package main

type visitor interface {
	visit(interface{})
}

type visitorMethod interface {
	visitFile(*file)
	visitDirectory(*directory)
}

func switchMethod(v visitor, itf interface{}) {
	vm, ok := v.(visitorMethod)
	if !ok {
		panic("v do not implement visitorMethod interface")
	}
	switch itf.(type) {
	case *file:
		f := itf.(*file)
		vm.visitFile(f)
	case *directory:
		d := itf.(*directory)
		vm.visitDirectory(d)
	default:
		panic("itf is not an entry object")
	}
}
