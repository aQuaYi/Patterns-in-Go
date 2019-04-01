package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func read3(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func main() {
	fmt.Println(read3("big7.txt"))
}
