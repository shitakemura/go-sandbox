package main

import (
	"fmt"
)

func main() {
	var m0 map[int]bool
	fmt.Printf("%#v, %t\n", m0, m0 == nil)

	m1 := map[int]bool{}
	fmt.Printf("%#v, %t\n", m1, m1 == nil)

	m2 := make(map[string]string)
	fmt.Printf("%#v, %t\n", m2, m2 == nil)

	m3 := map[string]int{
		"hello": 1,
		"world": 10,
	}
	fmt.Printf("%#v\n", m3)
}
