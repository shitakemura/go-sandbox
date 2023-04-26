package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("small")
	} else if n > 5 {
		fmt.Println("big")
	} else {
		fmt.Println("good")
	}
}
