package main

import "fmt"

func main() {
	var nilMap map[string]int

	fmt.Println(nilMap)
	fmt.Println(len(nilMap))
	fmt.Println(nilMap["abc"])

	totalWins := map[string]int{}
	fmt.Println(totalWins == nil)
	fmt.Println(totalWins["abc"])
	totalWins["abc"] = 3
	fmt.Println(totalWins["abc"])

	ages := make(map[string]int, 10)
	fmt.Println(ages, len(ages))

	fmt.Println("---")
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	delete(m, "hello")
	fmt.Println(m)
}
