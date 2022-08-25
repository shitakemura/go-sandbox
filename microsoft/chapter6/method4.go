package main

import "fmt"

type triangle struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

type coloredTriangle struct {
	triangle
	color string
}

func main() {
	t := coloredTriangle{triangle: triangle{size: 3}, color: "blue"}
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter:", t.perimeter())
	fmt.Println("Color:", t.color)
}
