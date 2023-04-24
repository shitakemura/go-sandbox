package main

import (
	"fmt"
)

func main() {
	var a [3]int

	a[1] = 1
	a[2] = 3
	fmt.Println(a)

	v := a[1]
	fmt.Println(v)
}