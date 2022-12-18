package main

import (
	"fmt"
	"math"
)

// find the length of the longest substring in it with no more than K distinct characters.
// Input: String="araaci", K=2
// Output: 4
// Explanation: The longest substring with no more than '2' distinct characters is "araa".

// Input: String="araaci", K=1
// Output: 2
// Explanation: The longest substring with no more than '1' distinct characters is "aa".
func main() {
	str := "araaci"
	k := 2

	mapRes := make(map[rune]int)
	slidingStart := 0
	maxlen := 0
	for index, key := range str { //a r a a c i
		_, flag := mapRes[key]
		if flag {
			mapRes[key] = 1
		} else {
			mapRes[key] += 1
		}
		for len(mapRes) > k { // a r a a c
			StartChar := str[slidingStart]
			mapRes[rune(StartChar)] -= 1 //decrement the value of currentchar in a map
			//count of current char values corresponding to key
			if mapRes[rune(StartChar)] == 0 {
				delete(mapRes, rune(StartChar))
			}
			slidingStart += 1 //shorten the sliding window
		}
		maxlen = int(math.Max(float64(maxlen), float64(index)-float64(slidingStart)+1)) //4
	}
	fmt.Println("maxlen ", maxlen)

}
