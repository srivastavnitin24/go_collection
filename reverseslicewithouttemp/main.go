package main

import (
	"fmt"
	"sort"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	x := []int{3, 2, 1, 10, 0, 89}
	y := []string{"a", "d", "c"}
	sort.Sort(sort.Reverse(sort.IntSlice(x)))
	sort.Strings(y)
	fmt.Println("=============== : ", y)

	//with temp variable
	reverse(x)
	fmt.Println(x)

	reverseCity()
}

func reverseCity() {

	cities := []string{"New York", "Beijing", "Bangkok", "Adelaide", "Tokyo", "Seoul", "Zurich", "Seattle"}

	fmt.Println("Before : ", cities)
	reversed := []string{}
	for i := range cities {
		n := cities[len(cities)-1-i]
		//fmt.Println(n) -- sanity check
		reversed = append(reversed, n)
	}

	fmt.Println("After : ", reversed)

}
