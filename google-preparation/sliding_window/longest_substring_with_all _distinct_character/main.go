package main

import (
	"fmt"
	"math"
)


//Given a string, find the length of the longest substring, which has all distinct characters.

func main() {
	str := "aabccbb"
	maxlen := 0
	slidingStart := 0
	m := make(map[rune]int)
	for index, key := range str {
		_, ok := m[key]
		if ok {
			slidingStart = int(math.Max(float64(slidingStart), float64(m[rune(str[index])]+1)))
		} else {
			m[key] = index
		}
		maxlen = int(math.Max(float64(maxlen), float64(index-slidingStart+1)))
	}
	fmt.Println("maxlen : ", maxlen)
}
