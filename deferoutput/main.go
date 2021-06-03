package main

import "fmt"

func main() {
	first()
	fmt.Println("Returned normally from first.")
	//if r := recover(); r != nil {
	//	fmt.Println("recoverd")
	//}
}

func first() {
	fmt.Println("Calling second.")
	second(0)
	fmt.Println("Returned normally from second.")
}

func second(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))

	}
	defer fmt.Println("Defer in second", i)
	//defer recoverSecond()
	fmt.Println("Printing in second", i)
	second(i + 1)
}

func recoverSecond() {
	if r := recover(); r != nil {
		fmt.Println("Nitin Srivastav", r)
	}
}

//=========================
//Returned normally from first.
//Calling second.
//Printing in second 0
//Printing in second 1
//Printing in second 2
//Printing in second 3
//Panicking!
//Defer in second 3
//Defer in second 2
//Defer in second 1
//Defer in second 0
//panic at 4
