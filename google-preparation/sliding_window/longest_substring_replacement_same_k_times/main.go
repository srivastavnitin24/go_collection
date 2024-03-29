package main

import (
	"fmt"
	"math"
)

/*
Example 1:

Input: String="aabccbb", k=2
Output: 5
Explanation: Replace the two 'c' with 'b' to have the longest repeating substring "bbbbb".
Example 2:

Input: String="abbcb", k=1
Output: 4
Explanation: Replace the 'c' with 'b' to have the longest repeating substring "bbbb".
Example 3:

Input: String="abccde", k=1
Output: 3
Explanation: Replace the 'b' or 'd' with 'c' to have the longest repeating substring "ccc".


/*
Given a string with lowercase letters only, if you are allowed to replace no more than k letters with any letter,
find the length of the longest substring having the same letters after replacement.

This problem follows the Sliding Window pattern, and we can use a similar dynamic sliding window strategy as discussed in Longest Substring with Distinct Characters. We can use a HashMap to count the frequency of each letter.

We will iterate through the string to add one letter at a time in the window.
We will also keep track of the count of the maximum repeating letter in any window (let’s call it maxRepeatLetterCount).
So, at any time, we know that we do have a window with one letter repeating maxRepeatLetterCount times; this means we should try to replace the remaining letters.
If the remaining letters are less than or equal to k, we can replace them all.
If we have more than k remaining letters, we should shrink the window as we cannot replace more than k letters.
While shrinking the window, we don’t need to update maxRepeatLetterCount. Since we are only interested in the longest valid substring, our sliding windows do not have to shrink, even if a window may cover an invalid substring. Either we expand the window by appending a character to the right or we shift the entire window to the right by one. We only expand the window when the frequency of the newly added character exceeds the historical maximum frequency (from a previous window that included a valid substring). In other words, we do not need to know the exact maximum count of the current window. The only thing we need to know is whether the maximum count exceeds the historical maximum count, and that can only happen because of the newly added char.
*/
func main() {
	str := "aabccbb"
	k := 2
	m := make(map[rune]int)
	maxCountRepeatingElement := 0
	slidingStart := 0
	maxlen := 0
	for index, char := range str {
		_, ok := m[char]
		if ok {
			m[char] += 1
		} else {
			m[char] = 1
		}
		maxCountRepeatingElement = int(math.Max(float64(maxCountRepeatingElement), float64(m[char])))
		if (index - slidingStart + 1 - maxCountRepeatingElement) > k {
			currentCh := m[rune(slidingStart)]
			_, ok := m[rune(currentCh)]
			if ok {
				m[rune(currentCh)] -= 1
			}
			slidingStart += 1
		}
		maxlen = int(math.Max(float64(maxlen), float64(index-slidingStart+1)))
	}
	fmt.Println("maxlen : ", maxlen)
}
