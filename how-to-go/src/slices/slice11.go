package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	rand.Seed(time.Now().UnixNano())

	fmt.Println(s[rand.Intn(len(s))])
}
