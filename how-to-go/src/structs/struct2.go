package main

import (
	"fmt"
)

type S struct {
	V1 int
	V2 string
}

func main() {
	var s0 S
	fmt.Printf("%#v\n", s0)

	s1 := S{
		V1: 100,
		V2: "aaa",
	}
	fmt.Printf("%#v\n", s1)

	s2 := &S{
		V1: 100,
		V2: "aaa",
	}
	fmt.Printf("%#v\n", s2)

	s3 := new(S)
	fmt.Printf("%#v\n", s3)
}