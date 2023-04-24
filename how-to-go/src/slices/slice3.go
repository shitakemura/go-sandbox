package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}

	for i, v := range s {
		fmt.Println(i, v)
	}
}
