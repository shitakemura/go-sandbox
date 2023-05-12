package main

import (
	"fmt"
	"sync"
)

type Test struct {
	name string
}

func main() {
	var wg sync.WaitGroup
	tests := []Test{Test{"a"}, Test{"b"}, Test{"c"}}

	for _, tt := range tests {
		wg.Add(1)
		go func(tt Test) {
			defer wg.Done()
			fmt.Println("1:", tt.name)
		}(tt)
	}

	wg.Wait()
	fmt.Println("done1")

	for i := range tests {
		v := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("2:", tests[v].name)
		}()
	}

	wg.Wait()
	fmt.Println("done2")
}
