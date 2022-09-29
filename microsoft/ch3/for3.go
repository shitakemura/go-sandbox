package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int32
	sec := time.Now().Unix()
	rand.Seed(sec)

	for {
		fmt.Printf("Writing inside the loop...")
		if num = rand.Int31n(10); num == 5 {
			fmt.Print(num)
			fmt.Println(" finish!")
			break
		}
		fmt.Println(num)
	}
}