package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{10, 11, 12}

	s3 := append(s1, s2...)
	fmt.Printf("%+v\n", s3)
}
