package main

import (
	"fmt"
)

type S struct{}

func (s *S) M() {
	fmt.Println("Hello World")
}

func main() {
	s := new(S)
	s.M()
}
