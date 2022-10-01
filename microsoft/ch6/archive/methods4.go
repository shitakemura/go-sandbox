package main

import "fmt"

type Triangle struct {
	size int
}

func (t *Triangle) doubleSize() {
	t.size *= 2
}

func (t *Triangle) SetSize(size int) {
	t.size = size
}

func (t *Triangle) Perimeter() int {
	t.doubleSize()
	return t.size * 3
}

func main() {
	t := Triangle{}
	t.SetSize(3)
	fmt.Println("Perimeter", t.Perimeter())
}
