package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("./dat")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open("./dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	_, err = f.Seek(6, 0)
	check(err)

	b2 := make([]byte, 3)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes: %s\n", n2, string(b2[:n2]))
}
