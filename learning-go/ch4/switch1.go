package main

import "fmt"

func main() {
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is short")
		case wordLen > 10:
			fmt.Println(word, "is long")
		default:
			fmt.Println(word, "is good")
		}
	}
}
