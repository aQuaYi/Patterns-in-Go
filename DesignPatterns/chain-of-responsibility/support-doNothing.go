package main

type doNothing struct {
	*baseSupport
}

func newDoNothing(name string) *doNothing {
	dn := &doNothing{
		baseSupport: newSupport(name),
	}
	dn.resolve = func(t *trouble) bool {
		return false
	}
	return dn
}
