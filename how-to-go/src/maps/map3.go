package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"aaa": 10,
		"bbb": 20,
	}

	v, ok := m["aaa"]
	if ok {
		fmt.Println(v)
	}

	if v, ok := m["ccc"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("ccc is not found")
	}
}
