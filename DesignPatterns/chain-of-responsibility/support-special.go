package main

type specialSupport struct {
	*baseSupport
	num int
}

func newSpecialSupport(name string, num int) *specialSupport {
	ss := &specialSupport{
		baseSupport: newSupport(name),
		num:         num,
	}
	ss.resolve = func(t *trouble) bool {
		if t.getNumber() == num {
			return true
		}
		return false
	}
	return ss
}
