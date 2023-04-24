package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []int{2, 3, 6, 8, 1, 5, 4, 7, 9}
	sort.Ints(s1)
	fmt.Println(s1)

	s2 := []float64{3.1, 1.1, 0.01, -5.25}
	sort.Float64s(s2)
	fmt.Println(s2)

	s3 := []string{"apple", "cat", "orange", "elephant"}
	sort.Slice(s3, func(i, j int) bool {
		return len(s3[i]) > len(s3[j])
	})
	fmt.Println(s3)
}
