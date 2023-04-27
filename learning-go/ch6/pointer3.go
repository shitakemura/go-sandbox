package main

import "fmt"

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func main() {
	var f *int
	fmt.Println(f)
	failedUpdate(f)
	fmt.Println(f)
}