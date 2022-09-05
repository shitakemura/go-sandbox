package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	ints2 := []int{1, 2, 3}
	s := sort.IntsAreSorted(ints2)
	fmt.Println("Sorted:", s)
}