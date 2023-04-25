package main

import (
	"errors"
	"fmt"
)

func main() {
	if err := func1(); err != nil {
		fmt.Println(err)
	}

	if err := func2(); err != nil {
		fmt.Println(err)
	}
}

func func1() error {
	return errors.New("Hello New Error")
}

func func2() error {
	return fmt.Errorf("Hello %s Error", "fmt")
}
