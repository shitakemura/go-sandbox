package main

import (
	"fmt"
)

func main() {
	m := map[string]int{}

	m["aaa"] = 10
	m["bbb"] = 20
	m["ccc"] = 30
	fmt.Println(m)

	m["bbb"] = 21
	fmt.Println(m)

	delete(m, "ccc")
	fmt.Println(m)
}
