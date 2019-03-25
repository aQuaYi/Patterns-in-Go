package main

type oddSupport struct {
	*baseSupport
}

func newOddSupport(name string) *oddSupport {
	os := &oddSupport{
		baseSupport: newSupport(name),
	}
	os.resolve = func(t *trouble) bool {
		if t.getNumber()%2 == 1 {
			return true
		}
		return false
	}
	return os
}
