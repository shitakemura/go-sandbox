package main

import (
	"ch9/formatter"
	"ch9/math"
	"fmt"
)

func main() {
	num := math.Double(2)
	output := formatter.Format(num)
	fmt.Println(output)
}
