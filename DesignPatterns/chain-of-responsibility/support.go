package main

import "fmt"

type support struct {
	name    string
	next    *support
	resolve func(*trouble) bool
}

func newSupport(name string) *support {
	return &support{
		name: name,
	}
}

func (s *support) setNext(ns *support) {
	s.next = ns
}

func (s *support) support(t *trouble) {
	if s.resolve(t) {
		s.done(t)
	} else if s.next != nil {
		s.next.support(t)
	} else {
		s.fail(t)
	}
}

func (s *support) done(t *trouble) {
	fmt.Printf("%s is resolved by %s.\n", t, s)
}

func (s *support) fail(t *trouble) {
	fmt.Printf("%s can NOT be resolved.\n", t)
}

func (s *support) String() string {
	return fmt.Sprintf("[%s]", s.name)
}
