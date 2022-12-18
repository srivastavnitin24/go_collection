package main

import "fmt"

/*
Given a string and a pattern, find out if the string contains any permutation of the pattern.
Permutation is defined as the re-arranging of the characters of the string. For example, “abc” has the following six permutations:
abc
acb
bac
bca
cab
cba

Example 1:

Input: String="oidbcaf", Pattern="abc"
Output: true
Explanation: The string contains "bca" which is a permutation of the given pattern.
Example 2:

Input: String="odicf", Pattern="dc"
Output: false
Explanation: No permutation of the pattern is present in the given string as a substring.
Example 3:

Input: String="bcdxabcdy", Pattern="bcdyabcdx"
Output: true
Explanation: Both the string and the pattern are a permutation of each other.
Example 4:

Input: String="aaacb", Pattern="abc"
Output: true
Explanation: The string contains "acb" which is a permutation of the given pattern.

*/
func main() {
	str := "abcd"
	pattern := "abc"
	matched := 0
	slidingStart := 0
	mapPattern := make(map[rune]int, 0)
	//insert pattern in a map
	for _, key := range pattern {
		_, ok := mapPattern[key]
		if ok {
			mapPattern[key] += 1
		} else {
			mapPattern[key] = 1
		}
	}

	for index, key := range str {
		_, ok := mapPattern[key]
		if ok {
			mapPattern[key] -= 1
			if mapPattern[key] == 0 {
				matched += 1
			}
		}
		if matched == len(mapPattern) {
			fmt.Println("True")
		}

		for index >= len(pattern)-1 {
			slidingStart += 1
			mapPattern[rune(slidingStart)] += 1
		}
	}

}
