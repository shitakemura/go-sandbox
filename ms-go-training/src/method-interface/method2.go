package main

import "fmt"

type triangle struct {
	size int
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

func main() {
	t := triangle{3}
	t.doubleSize()
	fmt.Println("Size:", t.size)
}
