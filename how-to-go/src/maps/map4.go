package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"aaa": 10,
		"bbb": 20,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
