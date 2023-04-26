package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "can divide 3 and 5")
			continue
		}
		if i%3 == 0 {
			fmt.Println(i, "can divide 3")
			continue
		}
		if i%5 == 0 {
			fmt.Println(i, "can divide 5")
			continue
		}
		fmt.Println(i)
	}
}
