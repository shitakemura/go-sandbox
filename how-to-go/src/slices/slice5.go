package main

import (
	"fmt"
)

func main() {
	s := make([][]int, 2)
	for i := 0; i < len(s); i++ {
		s[i] = make([]int, 3)
	}
	fmt.Printf("%#v\n", s)
}
