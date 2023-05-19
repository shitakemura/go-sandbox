package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		fmt.Println("Timer 2 fired")
		<-timer2.C
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("done")
}
