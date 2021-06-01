package main

import "fmt"

func main() {
	fmt.Println("go routines before")
	chan1 := make(chan int)
	go xyz(chan1)
	for i2 := range chan1 {
		fmt.Println(" ------- ", i2)
	}
	fmt.Println("go routines after ")

}
func xyz(chan1 chan int) chan int {
	for i := 0; i < 10; i++ {
		chan1 <- i
	}
	return chan1
}
