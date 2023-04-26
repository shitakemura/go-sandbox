package main

import "fmt"

func main() {
	x := make([]int, 5)
	fmt.Println(x)

	y := make([]int, 0, 10)
	y = append(y, 5, 6, 7, 8)
	fmt.Println(y, len(y), cap(y))
}
