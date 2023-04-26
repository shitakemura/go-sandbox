package main

import "fmt"

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "good bye"
	delete(m, 1)
}

func modSlice(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
	s = append(s, 10)
}

func main() {
	m := map[int]string{
		1: "１番目",
		2: "２番目",
	}
	fmt.Println(m)
	modMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	fmt.Println(s)
	modSlice(s)
	fmt.Println(s)
}