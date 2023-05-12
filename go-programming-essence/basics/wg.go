package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // リファレンスカウントを+1する
	
	go func() {
		defer wg.Done() // リファレンスカウントを-1する

		fmt.Println("重たい処理を実行")
	}()

	fmt.Println("別の重たい処理を実行")
	wg.Wait() // リファレンスカウンタが0になるまで待つ
	fmt.Println("done")
}