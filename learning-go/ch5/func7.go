package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println("＊初期データ")
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println("＊LastNameでソート")
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("＊Ageでソート")
	fmt.Println(people)
}
