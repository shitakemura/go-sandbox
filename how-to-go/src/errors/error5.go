package main

import (
	"errors"
	"fmt"
)

type MyError1 struct {
	Number  int
	Message string
}

func (e *MyError1) Error() string {
	return fmt.Sprintf("[%d] %s", e.Number, e.Message)
}

type MyError2 string

func (e *MyError2) Error() string {
	return string(*e)
}

func main() {
	var myError1 *MyError1
	var myError2 *MyError2

	e1 := &MyError1{Number: 1, Message: "first error"}
	fmt.Println(errors.As(e1, &myError1))
	fmt.Println(errors.As(e1, &myError2))

	e2 := MyError2("second error")
	fmt.Println(errors.As(&e2, &myError1))
	fmt.Println(errors.As(&e2, &myError2))

	e3 := fmt.Errorf("error: %w", &MyError1{Number: 3, Message: "third error"})
	fmt.Println(errors.As(e3, &myError1))
	fmt.Println(errors.As(e3, &myError2))
}
