package main

import "fmt"

type S1 struct{}

func (s *S1) String() string {
	return "S1.String"
}

type S2 struct{}

func (s *S2) String() string {
	return "S2.String"
}

type I interface {
	String() string
}

func main() {
	s1 := new(S1)
	s2 := new(S2)

	var i1 I = s1
	var i2 I = s2

	s1_2, ok := i1.(*S1)
	fmt.Printf("%#v, %v\n", s1_2, ok)

	s2_2, ok := i1.(*S2)
	fmt.Printf("%#v, %v\n", s2_2, ok)

	switch v := i2.(type) {
	case *S2:
		fmt.Printf("%#v\n", v)
	}
}
