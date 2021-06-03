package main

import (
	"fmt"
	"sync"
)

var count =0

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go dosomething(&wg)
	}
	wg.Wait()
	fmt.Println("====== ", count)
}

func dosomething(wg *sync.WaitGroup) {
	count++
	wg.Done()
}
