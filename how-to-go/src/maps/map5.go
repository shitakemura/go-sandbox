package main

import (
	"fmt"
)

func main() {
	m := map[string]map[int]int{
		"aaa": map[int]int{
			1: 2,
			3: 4,
		},
		"bbb": map[int]int{
			5: 6,
			7: 8,
		},
	}
	fmt.Println(m)
}
