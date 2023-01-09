package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	fmt.Println(v)
	ScaleFunc(&v, 2)
	fmt.Println(v)

	p := &Vertex{3, 4}
	p.Scale(2)
	fmt.Println(p)
	ScaleFunc(p, 2)
	fmt.Println(p)
}
