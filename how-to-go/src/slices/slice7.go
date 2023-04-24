package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := s[1:3]
	fmt.Println(s1)

	s2 := s[:2]
	fmt.Println(s2)

	s3 := s[2:]
	fmt.Println(s3)

	s4 := s[5:]
	fmt.Println(s4)

	s4[0] = 20
	s4[1] = 31
	fmt.Println(s4)
	fmt.Println(s)
}
