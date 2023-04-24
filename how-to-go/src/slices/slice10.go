package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v len:%d, cap:%d\n", s, len(s), cap(s))

	i := 2
	s1 := s[:i]
	s2 := s[i+1:]
	s = append(s1, s2...)
	fmt.Printf("%v len:%d, cap:%d\n", s, len(s), cap(s))

	s = []int{1, 2, 3, 4, 5}
	s = append(s[:i], s[i+1:]...)
	fmt.Printf("%v len:%d, cap:%d\n", s, len(s), cap(s))
}
