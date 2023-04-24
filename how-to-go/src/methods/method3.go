package main

import (
	"fmt"
)

type S1 struct{}

func (s *S1) M() string {
	return "from S1.M"
}

type S2 struct {
	*S1
}

func (s *S2) M() string {
	return "from S2.M"
}

func main() {
	s := new(S2)
	fmt.Println(s.M())
	fmt.Println(s.S1.M())
}