package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("sum: %d, lastUpdated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("NG: ", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("OK: ", c.String())
}

func main() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

	fmt.Println("---")

	var c1 Counter
	fmt.Println(c1.String())
	doUpdateWrong(c1)
	fmt.Println(c1.String())
	doUpdateRight(&c1)
	fmt.Println(c1.String())
}
