package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I love india"
	fmt.Println("Before reverse : ", str)
	s := reverse1(str)
	fmt.Println("After reverse1 : ", s)

	s2 := reverse2(str)
	fmt.Println("After reverse2 : ", s2)
}

func reverse1(str string) string {
	result := ""

	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func reverse2(str string) string {
	res := ""
	newstr := strings.Split(str, " ")
	for _, j := range newstr {
		result := ""
		for _, v := range j {
			result = string(v) + result
		}
		res = res + " " + result
	}

	return res
}
