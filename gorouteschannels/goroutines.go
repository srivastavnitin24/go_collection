package gorouteschannels

import "fmt"

func SumOfSquares(c, quit chan int) {
	y := 1
	for {
		select {
		case c <- (y * y):
			y++
		case <-quit:
			return
		}
	}
}

func mainfunc() int {
	mychannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0

	go func() {
		for i := 1; i <= 5; i++ {
			sum += <-mychannel
		}
		fmt.Println(sum)

		quitchannel <- 0
	}()
	fmt.Println("before SumOfSquares")
	SumOfSquares(mychannel, quitchannel)

	return sum

}