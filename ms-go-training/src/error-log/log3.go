package main

import (
	"fmt"
	"log"
)

func main() {
	log.Panic("Hey, I'm an error log!")
	fmt.Printf("Can you see me?")
}