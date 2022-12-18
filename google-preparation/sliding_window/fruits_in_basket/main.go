package main

import (
	"fmt"
	"math"
)

var maxNum = 0

// You are visiting a farm to collect fruits. The farm has a single row of fruit trees. You will be given two baskets, and your goal is to pick as many fruits as possible to be placed in the given baskets.

// You will be given an array of characters where each character represents a fruit tree. The farm has following restrictions:

// Each basket can have only one type of fruit. There is no limit to how many fruit a basket can hold.
// You can start with any tree, but you can’t skip a tree once you have started.
// You will pick exactly one fruit from every tree until you cannot, i.e., you will stop when you have to pick from a third fruit type.
// Write a function to return the maximum number of fruits in both baskets.

// Example 1:
// Input: Fruit=['A', 'B', 'C', 'A', 'C']
// Output: 3
// Explanation: We can put 2 'C' in one basket and one 'A' in the other from the subarray ['C', 'A', 'C']

// Example 2:
// Input: Fruit=['A', 'B', 'C', 'B', 'B', 'C']
// Output: 5
// Explanation: We can put 3 'B' in one basket and two 'C' in the other basket.
// This can be done if we start with the second letter: ['B', 'C', 'B', 'B', 'C']

func main() {
	s := []rune{'A', 'B', 'C', 'A', 'C'}
	slidingStart := 0

	maps := make(map[rune]int)
	for index, key := range s {
		_, flag := maps[key]
		if flag {
			maps[key] = 1
		} else {
			maps[key] += 1
		}
		for len(maps) > 2 {
			currentChar := s[slidingStart]
			maps[currentChar] -= 1      //decrement the value of currentchar in a map
			if maps[currentChar] == 0 { //count of currentchar values corresponding to key
				delete(maps, rune(currentChar))
			}
			slidingStart += 1 //shorten the sliding window
		}
		maxNum = int(math.Max(float64(maxNum), float64(index-slidingStart+1)))
	}
	fmt.Println("result : ", maxNum)
}
