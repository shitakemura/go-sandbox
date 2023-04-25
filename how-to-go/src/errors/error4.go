package main

import (
	"errors"
	"fmt"
)

type MyError string

func (e MyError) Error() string {
	return string(e)
}

const MyError400 MyError = "not found"
const MyError500 MyError = "internal server error"

func main() {
	e1 := MyError400
	fmt.Println(errors.Is(e1, MyError400))
	fmt.Println(errors.Is(e1, MyError500))

	e2 := MyError500
	fmt.Println(errors.Is(e2, MyError400))
	fmt.Println(errors.Is(e2, MyError500))

	e3 := fmt.Errorf("error: %w", MyError400)
	fmt.Println(errors.Is(e3, MyError400))
	fmt.Println(errors.Is(e3, MyError500))
}
