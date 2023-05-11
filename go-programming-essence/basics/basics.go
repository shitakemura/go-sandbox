package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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

func showName(user *User) {
	fmt.Println(user.name)
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

type Value int

func (v Value) Add(n Value) Value {
	return v + n
}

func (v *Value) Add2(n Value) {
	*v += n
}

func printDetail(v interface{}) {
	switch t := v.(type) {
	case int, int32, int64:
		fmt.Println("int/int32/int64 type:", t)
	case string:
		fmt.Println("string type:", t)
	default:
		fmt.Println("unknown type")
	}
}

// constructor
type User2 struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User2 {
	return &User2{
		Name: name,
		Age:  age,
	}
}

// interface
type Speaker interface {
	Speak() error
}

type Dog struct{}

func (d *Dog) Speak() error {
	fmt.Println("BowWow")
	return nil
}

type Cat struct{}

func (c *Cat) Speak() error {
	fmt.Println("Meow")
	return nil
}

func DoSpeak(s Speaker) error {
	return s.Speak()
}

// defer

func doSomething(dir string) error {
	err := os.Mkdir(dir, 0755)
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	f, err := os.Create(filepath.Join(dir, "data.txt"))
	if err != nil {
		return err
	}
	defer f.Close()

	// ファイルを使った処理
	return nil
}

func doSomethingAndDefer1() {
	var n = 1
	defer func() {
		fmt.Println(n)
	}()

	n = 2
}

func doSomethingAndDefer2() {
	var n = 1
	defer fmt.Println(n)

	n = 2
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

	// type
	type MyString string
	var ms MyString = "foo"
	fmt.Println(ms)

	// struct
	usr := User{name: "Bob"}
	showName(&usr)

	// method
	v2 := Value(1)
	v2 = v2.Add(2)
	fmt.Println(v2)

	// pointer
	v3 := 1
	p := &v3
	fmt.Println(p)

	user2 := new(User) // make pointer
	user2.name = "taro"

	v4 := Value(1)
	v4.Add2(2)
	fmt.Println(v4)

	// interface{}, any
	var v5 interface{}

	v5 = 1
	n5 := v5.(int)
	fmt.Println(n5)

	v5 = "hello world"
	s5 := v5.(string)
	fmt.Println(s5)

	s5, ok = v5.(string)
	if !ok {
		fmt.Println("v5 is not string type")
	} else {
		fmt.Printf("v is %q\n", s5)
	}

	type V int
	var v6 V = 123
	printDetail(v6)

	// constructor
	u2 := NewUser("Taro", 100)
	fmt.Println(u2)

	// interface
	dog := Dog{}
	dog.Speak()
	DoSpeak(&dog)

	cat := &Cat{}
	cat.Speak()
	DoSpeak(cat)

	// defer
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var b [512]byte
	nn, err := f.Read(b[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b[:nn]))

	doSomethingAndDefer1()
	doSomethingAndDefer2()
}
