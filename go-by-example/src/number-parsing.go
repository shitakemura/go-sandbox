package main

import (
	"fmt"
	"strconv"
)

func main() {
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}