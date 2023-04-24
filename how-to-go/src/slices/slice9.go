package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 2, 1, 4, 1, 3}

	m := map[int]bool{}
	for _, v := range s {
		m[v] = true
	}

	u := []int{}
	for k, _ := range m {
		u = append(u, k)
	}

	fmt.Println(u)
}
