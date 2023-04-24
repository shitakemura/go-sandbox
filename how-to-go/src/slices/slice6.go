package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := []int{}

	for _, v := range s1 {
		if v%2 == 0 {
			s2 = append(s2, v)
		}
	}
	fmt.Printf("%+v\n", s2)
}
