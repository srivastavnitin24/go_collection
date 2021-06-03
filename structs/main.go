package main

import (
	"fmt"
)

type Person struct {
	personname string
	Address
}

type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	a1 := Address{"Pam", "Dehradun", 2200}

	a2 := Address{Name: "Ram", city: "Delhi", Pincode: 2400}

	a3 := Address{Name: "Sam", city: "Lucknow", Pincode: 1070}

	//mp := make(map[int]Address)

	sample := map[int]Address{1: a1, 2: a2, 3: a3}

	for _, v := range sample {
		//for _, v1 := range k {
		//fmt.Println(reflect.TypeOf(v))
		//fmt.Println(k, ":", v)
		if v.Pincode == 1070 {
			fmt.Println(v.Name)
		}
		//}

	}
}