package main

import "fmt"

func main() {
	numbers := [...]int{99: -1}
	fmt.Println("First position:", numbers[0])
	fmt.Println("Second position:", numbers[1])
	fmt.Println("Last position:", numbers[99])
	fmt.Println("Length:", len(numbers))
}