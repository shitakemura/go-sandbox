package main

import (
	"fmt"
)

type consumer func(int)

func provider(c consumer) {
	c(100)
}

func main() {
	printer := func(n int) {
		fmt.Printf("printer: %d\n", n)
	}
	provider(printer)

	total := 0
	aggregator := func(n int) {
		total += n
	}
	provider(aggregator)
	provider(aggregator)
	provider(aggregator)
	fmt.Println(total)

}
