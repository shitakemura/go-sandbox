package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}

	v := s[0]
	fmt.Println(v)

	s[1] = 10
	fmt.Printf("%v\n", s)

	s = append(s, 11)
	s = append(s, 12)
	fmt.Printf("%v\n", s)
}
