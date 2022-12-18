package main

import "fmt"

/**
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
Example 1:
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:
Input: "cbbd"
Output: "bb"
*/
func main() {

	str := []byte("abad")
	length := len(str)
	if length == 0 {
		fmt.Println("string in length zero")
	}

	dp := make([][]bool, length)
	//[[] [] [] []]
	for i := range dp {
		dp[i] = make([]bool, length)
	}
	//[[false false false false] [false false false false] [false false false false] [false false false false]]

	result := []byte{}
	max := 0
	for i := length - 1; i >= 0; i-- { // traversing the string in reverse from 3->0
		for j := i; j < length; j++ { //3 -> 2 ->
			if str[i] == str[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if max < j-i+1 {
					max = j - i + 1
					result = str[i : j+1]
				}
			}
		}
	}
	fmt.Println("--------- : ", string(result))
}
