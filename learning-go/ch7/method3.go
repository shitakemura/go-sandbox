package main

import "fmt"

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	myAdder10 := Adder{start: 10} // method value
	fmt.Println(myAdder10.AddTo(5))

	f1 := myAdder10.AddTo
	fmt.Println(f1(10))

	f2 := Adder.AddTo // method expression: func(a Adder, val int) int
	fmt.Println(f2(myAdder10, 15))
}
