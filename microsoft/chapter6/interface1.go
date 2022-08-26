package main

import "fmt"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	size float64
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func main() {
	var s Shape = Square{3}
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}
