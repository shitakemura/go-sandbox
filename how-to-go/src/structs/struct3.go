package main

import (
	"fmt"
)

type S struct {
	V1 int
	V2 string
}

func main() {
	s1 := S{V1: 100, V2: "aaa"}

	v1 := s1.V1
	fmt.Println(v1)

	s1.V2 = "bbb"
	fmt.Printf("%#v\n", s1)
}
