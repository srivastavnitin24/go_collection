package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
	if r := recover(); r != nil {
		fmt.Println("recoverd")
	}
}

func f() {
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))

	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}



//=========================
//Returned normally from f.
//Calling g.
//Printing in g 0
//Printing in g 1
//Printing in g 2
//Printing in g 3
//Panicking!
//Defer in g 3
//Defer in g 2
//Defer in g 1
//Defer in g 0
//panic at 4
