package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

var (
	once     sync.Once
	instance *singleton
)

// thread safe
func newSingleton() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	s1 := newSingleton()
	s2 := newSingleton()

	fmt.Println(s1 == s2)
}
