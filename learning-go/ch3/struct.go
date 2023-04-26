package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		pet  string
	}

	bob := person{}
	fmt.Println(bob)

	taro := person{"taro", 8, "jiro"}
	fmt.Println(taro)

	hanako := person{name: "hanako", age: 3}
	fmt.Println(hanako)

	tom := person{
		name: "",
		age:  0,
		pet:  "",
	}
	fmt.Println(tom)
}
