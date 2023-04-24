package main

import (
	"fmt"
)

type S1 struct {
	V1 int
	V2 string
}

type s2 struct {
	V1 bool
	V2 string
}

func main() {
	fmt.Printf("%#v\n", S1{})
	fmt.Printf("%#v\n", s2{})
}
