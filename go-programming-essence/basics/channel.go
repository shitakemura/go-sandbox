package main

import (
	"fmt"
	"sync"
)

// channel: goroutineに対してデータを送受信する仕組み

func server(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	defer close(ch)

	ch <- "one"
	ch <- "two"
	ch <- "three"
}

func main() {
	var s string
	var wg sync.WaitGroup

	ch := make(chan string)

	wg.Add(1)
	go server(&wg, ch)

	s = <-ch
	fmt.Println(s)

	s = <-ch
	fmt.Println(s)

	s = <-ch
	fmt.Println(s)

	wg.Wait()

	ch = make(chan string)
	wg.Add(1)
	go server(&wg, ch)

	for s := range ch {
		fmt.Println(s)
	}

	wg.Wait()
}
