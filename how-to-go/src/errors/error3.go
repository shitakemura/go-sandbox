package main

import (
	"fmt"
)

func func1() error {
	return fmt.Errorf("func1 error:\n%w", func2())
}

func func2() error {
	return fmt.Errorf("func2 error:\n%w", func3())
}

func func3() error {
	return fmt.Errorf("func3 error")
}

func main() {
	e := func1()
	fmt.Println(e)
}