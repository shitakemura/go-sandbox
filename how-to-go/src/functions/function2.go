package main

import (
	"fmt"
)

func main() {
	f := func(s string) int {
		return len(s)
	}

	fmt.Println(f("hello world"))

	fmt.Println(funcStringIntUser(f))
}

func funcStringIntUser(f func(string) int) string {
	return fmt.Sprintf("function return value is %d", f("test"))
}
