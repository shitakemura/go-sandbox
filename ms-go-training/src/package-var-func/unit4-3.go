package main

import "fmt"

func updateName(name *string) {
	*name = "David"
}

func main() {
	firstName := "John"
	updateName(&firstName)
	fmt.Println(firstName)
}
