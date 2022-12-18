package main

import (
	"fmt"
)

type Person struct {
	Name  string
	Likes []string
}

func main() {
	p1 := Person{
		Name:  "Nitin",
		Likes: []string{"A", "B"},
	}
	p2 := Person{
		Name:  "Ram",
		Likes: []string{"C", "D"},
	}

	var people []Person
	people = []Person{p1, p2}

	//likes := make(map[string][]Person)
	likes := make([]string, 0)
	for _, p := range people {
		fmt.Println(p.Name, " : ", p.Likes)
		for _, l := range p.Likes {
			//likes[l] = append(likes[l], p)
			likes = append(likes, l)
			fmt.Println(l)
		}
	}
	fmt.Println(likes)
}
