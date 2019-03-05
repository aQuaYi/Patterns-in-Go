package main

import "fmt"

func main() {
	size := 10
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Println(slice[i])
	}

	return
}

type iterator interface {
	hasNext() bool
	next() interface{}
}
