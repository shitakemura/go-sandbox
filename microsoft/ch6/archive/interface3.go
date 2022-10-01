package main

import "fmt"

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

func main() {
	rs := Person{Name: "John Doe", Country: "USA"}
	ab := Person{Name: "Mark Collins", Country: "United Kingdom"}
	fmt.Printf("%s\n%s\n", rs, ab)
}
