package main

import "fmt"

type S1 struct{}

func (s *S1) String() string {
	return "S1.String"
}

type S2 struct{}

type S3 struct{}

func (s *S3) String(in string) string {
	return "S3.String"
}

type I interface {
	String() string
}

func main() {
	s1 := new(S1)
	// s2 := new(S2)
	// s3 := new(S3)

	var i I

	i = s1
	fmt.Println(i)
	// i = s2
	// i = s3
}
