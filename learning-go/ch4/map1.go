package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 3,
		"c": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}
