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
	v := Vertex{X: 3, Y: 4}
	v.Scale(2)
	ScaleFunc(&v, 10)
	fmt.Println(v)

	p := &Vertex{X: 4, Y: 3}
	p.Scale(2)
	ScaleFunc(p, 10)
	fmt.Println(p)
}
