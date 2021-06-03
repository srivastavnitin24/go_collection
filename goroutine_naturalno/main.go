package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var number uint32 = 5
	var count uint32 = 1
	trigger := func(i uint32) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fmt.Println(i)
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(1 * time.Second)
		}
	}
	for i := uint32(1); i < number; i++ {
		go trigger(i)
	}
	trigger(number)
}



//package main
//
//import (
//	"fmt"
//	"sync/atomic"
//	"time"
//)
//
//func main() {
//	var number uint32 = 5
//	var count uint32
//	trigger := func(i uint32, fn func()) {
//		for {
//			if n := atomic.LoadUint32(&count); n == i {
//				fn()
//				// Must be added to the atom after the completion of the function
//				atomic.AddUint32(&count, 1)
//				break
//			}
//			time.Sleep(1 * time.Second)
//		}
//	}
//
//	for i := uint32(0); i < number; i++ {	//5
//		go func(i uint32) {
//			fn := func() {
//				fmt.Println(i)
//			}
//			trigger(i, fn)
//		}(i)
//	}
//	trigger(number, func() {})
//	// will print in natural order (must be like this)
//}
