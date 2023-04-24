package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello(3))
	fmt.Println(Echo("aaa", "bbb", "ccc"))
	fmt.Println(Double(100))
}

func Hello(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "hello "
	}
	return s
}

func Echo(in ...string) string {
	s := ""
	for _, v := range in {
		s += v + " "
	}
	return s
}

func Double(n int) (m int) {
	m = n * 2
	return
}
