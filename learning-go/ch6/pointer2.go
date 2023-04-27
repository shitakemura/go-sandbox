package main

import (
	"fmt"
)

func stringp(s string) *string {
	return &s
}

func main() {
	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	// s := "Perry"
	// p := person{
	// 	FirstName:  "Pat",
	// 	MiddleName: &s,
	// 	LastName:   "Peterson",
	// }

	p := person{
		FirstName:  "Pat",
		MiddleName: stringp("Perry"),
		LastName:   "Peterson",
	}

	fmt.Println(p)
	fmt.Println(*p.MiddleName)
}
