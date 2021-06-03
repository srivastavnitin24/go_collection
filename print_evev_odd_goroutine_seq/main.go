package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var number uint32 = 10
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				// Must be added to the atom after the completion of the function
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(1 * time.Millisecond)
		}
	}

	for i := uint32(0); i < number; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(number, func() {})
}

//==================================================================================
//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//func main() {
//	var wg sync.WaitGroup
//
//	flag := make(chan bool)
//	wg.Add(2)
//	go even(flag, &wg)
//	go odd(flag, &wg)
//	flag <- true
//	wg.Wait()
//}
//
//func even(chanEven chan bool, wg *sync.WaitGroup) {
//	for i := 2; i <= 50; i = i + 2 {
//		<-chanEven
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(i)
//
//		chanEven <- true
//	}
//	wg.Done()
//}
//func odd(chanOdd chan bool, wg *sync.WaitGroup) {
//	for i := 1; i <= 50; i = i + 2 {
//		<-chanOdd
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(i)
//
//		chanOdd <- true
//	}
//	wg.Done()
//}
