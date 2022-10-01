package main

import (
	"example/store"
	"fmt"
)

func main() {
	bruce, _ := store.CreateEmployee("Bruce", "Lee", 500)
	fmt.Println(bruce)
	fmt.Println(bruce.CheckCredits())

	credits, err := bruce.AddCredits(250)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance =", credits)
	}

	credits, err = bruce.RemoveCredits(2500)
	if err != nil {
		fmt.Println("Can't withdraw or overdrawn!", err)
	} else {
		fmt.Println("New Credits Balance =", credits)
	}

	bruce.ChangeName("Mark")
	fmt.Print(bruce)
}