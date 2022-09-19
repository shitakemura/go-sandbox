package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{i: 2, b: true},
		{i: 3, b: false},
		{i: 5, b: true},
		{i: 7, b: true},
		{i: 11, b: false},
		{i: 13, b: true},
	}
	fmt.Println(s)
}
