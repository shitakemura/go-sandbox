package main

import (
	"fmt"
)

func main() {
	pow := make([]int, 10)
	fmt.Println(pow)

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
		// pow[i] = int(math.Pow(2, float64(i)))
	}
	fmt.Println(pow)

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
