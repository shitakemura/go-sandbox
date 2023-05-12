package main

import "fmt"

type F struct {
	Name string
	Age int
}

// Stringer interfaceを実装
func (f *F) String() string {
	return fmt.Sprintf("Name=%q, Age=%d", f.Name, f.Age)
}

func main() {
	f := &F{
		Name: "John",
		Age: 20,
	}

	fmt.Printf("%v\n", f)
}