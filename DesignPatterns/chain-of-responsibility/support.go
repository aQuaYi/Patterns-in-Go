package main

import "fmt"

type supporter interface {
	deal(t *trouble)
	setNext(supporter) supporter
}

type baseSupport struct {
	name    string
	next    supporter
	resolve func(*trouble) bool
}

func newSupport(name string) *baseSupport {
	return &baseSupport{
		name: name,
	}
}

func (s *baseSupport) setNext(si supporter) supporter {
	s.next = si
	return si
}

func (s *baseSupport) deal(t *trouble) {
	if s.resolve(t) {
		s.done(t)
	} else if s.next != nil {
		s.next.deal(t)
	} else {
		s.fail(t)
	}
}

func (s *baseSupport) done(t *trouble) {
	fmt.Printf("%s is resolved by %s.\n", t, s)
}

func (s *baseSupport) fail(t *trouble) {
	fmt.Printf("%s can NOT be resolved.\n", t)
}

func (s *baseSupport) String() string {
	return fmt.Sprintf("[%s]", s.name)
}
