package main

type limitSupport struct {
	*baseSupport
	limit int
}

func newLimitSupport(name string, limit int) *limitSupport {
	ls := &limitSupport{
		baseSupport: newSupport(name),
		limit:       limit,
	}
	ls.resolve = func(t *trouble) bool {
		if t.getNumber() < limit {
			return true
		}
		return false
	}
	return ls
}
