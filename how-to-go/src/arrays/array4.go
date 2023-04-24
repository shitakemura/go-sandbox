package main

import (
	"fmt"
)

func main() {
	a1 := [3]int{2, 4, 6}
	s1 := a1[:]
	fmt.Println(s1)

	s2 := []int{1, 3, 5}
	a2 := (*[3]int)(s2)
	fmt.Println(*a2)
}