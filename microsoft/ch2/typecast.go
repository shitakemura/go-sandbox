package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Printf("i - type: %T, value: %v\n", i, i)
	fmt.Printf("s - type: %T, value: %v\n", s, s)
}