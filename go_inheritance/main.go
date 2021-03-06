// Golang program to illustrate the
// concept of inheritance
package main

import (
	"fmt"
)

type Comic struct {
	Universe string
}

func (comic Comic) ComicUniverse() string {
	return comic.Universe
}

type Marvel struct {
	Comic
}

type DC struct {
	Comic
}

func main() {

	c1 := Marvel{
		Comic{
			Universe: "MCU",
		},
	}

	fmt.Println("Universe is:", c1.ComicUniverse())

	c2 := DC{
		Comic{
			Universe: "DC",
		},
	}

	fmt.Println("Universe is:", c2.ComicUniverse())
}
