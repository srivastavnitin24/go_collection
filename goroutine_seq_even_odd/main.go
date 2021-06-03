//Program a program to print natural number.

package main

import (
	"fmt"
	"sync"
	"time"
)

func odd(ch chan bool, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i = i + 2 {
		<-ch
		time.Sleep(1*time.Second)
		fmt.Println(i)
		ch <- true
	}
	wg.Done()
}

func even(ch chan bool, wg *sync.WaitGroup) {
	for i := 2; i <= 10; i = i + 2 {
		<-ch
		time.Sleep(1*time.Second)
		fmt.Println(i)
		ch <- true
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool,10)
	wg.Add(2)
	go even(ch, &wg)
	go odd(ch, &wg)
	ch <- true
	wg.Wait()
}
