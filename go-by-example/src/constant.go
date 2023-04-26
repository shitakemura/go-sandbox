package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Printf("type: %T\n", d)
	fmt.Println(int64(d))
	fmt.Printf("type: %T\n", int64(d))

	fmt.Println(math.Sin(n))
}
