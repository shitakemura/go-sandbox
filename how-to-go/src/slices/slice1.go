package main

import (
	"fmt"
)

func main() {
	var s0 []int
	fmt.Printf("%#v\n", s0)

	s1 := []int{}
	fmt.Printf("%v\n", s1)

	s2 := make([]int, 2)
	s3 := make([]string, 2, 4)
	fmt.Printf("%#v, %#v\n", s2, s3)
	fmt.Printf("s2: len:%d, cap:%d\n", len(s2), cap(s2))
	fmt.Printf("s3: len:%d, cap:%d\n", len(s3), cap(s3))

	s4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", s4)
}
