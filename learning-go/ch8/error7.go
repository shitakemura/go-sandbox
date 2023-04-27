package main

import (
	"errors"
	"fmt"
)

type MyErr struct {
	Message string
}

func (me MyErr) Error() string {
	return me.Message
}

func AFuntionThatReturnsAnError() error {
	return MyErr{Message: "my error"}
}

func main() {
	err := AFuntionThatReturnsAnError()
	var myErr MyErr
	if errors.As(err, &myErr) {
		fmt.Println(myErr.Message)
	}
}
