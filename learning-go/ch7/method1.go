package main

import "fmt"

type Person struct {
	LastName  string
	FirstName string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s:%s : Age: %d", p.LastName, p.FirstName, p.Age)
}

func main() {
	p := Person{
		LastName:  "hoge",
		FirstName: "fuga",
		Age:       20,
	}
	fmt.Println(p.String())
}
