package main

import (
	"fmt"
	"os"
)

type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2)

	i3, ok := i.(string)
	if !ok {
		err := fmt.Errorf("iの型が予想外です", i)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(i3)
}
