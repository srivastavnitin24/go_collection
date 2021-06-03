package main

import (
	"fmt"
	"time"
)

var count int = 1

func main() {
	number := 5
	for i := 0; i < number; i++ {
		go dosomething(i)
	}
	dosomething(number)

}

func dosomething(i int)  {
	for {
		if n := count; n == i {
			fmt.Println(i)
			count++
			break
		}
		time.Sleep(1 * time.Second)
	}
}
//===================================================================

//package main
//
//import (
//"fmt"
//"time"
//)
//
//func main() {
//	var count int32 = 1
//	var number int32 = 5
//
//	dosomething := func(i int32) {
//		for {
//			if n := atomic.LoadInt32(&count); n == i {
//				fmt.Println(i)
//				atomic.AddInt32(&count, 1)
//				break
//			}
//			time.Sleep(1 * time.Second)
//		}
//	}
//
//	for i := int32(0); i < number; i++ {
//		go dosomething(i)
//	}
//	dosomething(number)
//}