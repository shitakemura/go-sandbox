package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("./dat", d1, 0644)
	check(err)

	f, err := os.Create("./dat2")
	check(err)
	defer f.Close()
}
