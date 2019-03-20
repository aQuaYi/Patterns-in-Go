package main

type visitor interface {
	visit(entry)
}
