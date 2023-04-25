package main

import (
	"fmt"
)

type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return e.Message
}

func func1() error {
	return MyError{Message: "hello my error"}
}

func main() {
	if err := func1(); err != nil {
		fmt.Println(err)
	}
}
