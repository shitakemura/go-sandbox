package main

import (
	"fmt"
)

func main() {
	var a0 [3]int
	fmt.Printf("%#v\n", a0)

	a1 := [2]int{2, 3}
	fmt.Printf("%#v\n", a1)

	a2 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%#v\n", a2)
}
