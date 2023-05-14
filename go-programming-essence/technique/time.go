package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// time.Format / time.Parse
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format(time.RFC3339))

	var s = "2022/12/25 07:42:38"
	d, err := time.Parse("2006/01/02 15:04:05", s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)

	// time.Duration
	dd, err := time.ParseDuration("3s")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dd)

	dd, err = time.ParseDuration("4m")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dd)

	dd, err = time.ParseDuration("5h")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dd)

	fmt.Println(dd * 3)
}
