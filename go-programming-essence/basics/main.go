package main

import (
	"fmt"
	"log"
	"sort"
)

type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)

type User struct {
	name string
}

func findUserFromList(name string) (*User, error) {
	return &User{name}, nil
}

func FindUser(name string) (*User, error) {
	user, err := findUserFromList(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func userName() (User, error) {
	return User{"Taro"}, nil
}

func main() {
	// iota
	var fruit Fruit = Apple
	fmt.Println(fruit)

	// function
	user, err := FindUser("Bob")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.name)

	if user, err := userName(); err == nil {
		fmt.Println(user.name)
	} else {
		log.Println(err)
	}

	// switch
	i := 3
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("other")
	}

	hit := 10
	switch {
	case i > hit:
		fmt.Println("bigger than", hit)
	case i < hit:
		fmt.Println("less than", hit)
	case i == hit:
		fmt.Println("equal to", hit)
	}

	// slice
	a := []int{1, 2, 3}
	fmt.Println(a, len(a))
	a = append(a, 4, 5, 6)
	fmt.Println(a, len(a))

	a = make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

	// sliceの要素を削除
	// (1) 偶数のみ
	a2 := make([]int, 0, len(a)/2)
	for i := 0; i < len(a); i++ {
		if i%2 == 0 {
			a2 = append(a2, i)
		}
	}
	a = a2
	fmt.Println(a)

	// (2) 50を削除
	a = make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}
	n := 50
	a = append(a[:n], a[n+1:]...)
	fmt.Println(a)

	// map
	m := make(map[string]int)
	m["John"] = 21
	m["Bob"] = 18
	m["Mark"] = 33
	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	m = map[string]int{
		"Taro":   5,
		"Hanako": 20,
	}
	fmt.Println(m)
	delete(m, "Hanako")
	fmt.Println(m)

	// mapのソート
	m = map[string]int{
		"John": 21,
		"Bob":  18,
		"Mark": 33,
	}

	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Printf("key, %v, value: %v\n", k, m[k])
	}

	// no key
	m2 := map[string]string{
		"foo": "bar",
	}
	fmt.Println(m2["zoo"])

	v, ok := m["zoo"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("no key")
	}
}
