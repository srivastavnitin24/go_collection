package main

import "fmt"

func main() {
	var balInput string
	//fmt.Scan(&balInput)
	balInput = "[{()}]"
	res := BalancedBracketsCheck(balInput)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func BalancedBracketsCheck(brac string) bool {
	isBalanced := true
	bMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	s := make([]rune, 0, len(brac))

	for _, c := range brac {
		if _, ok := bMap[c]; ok == true {
			// new bracket will be appending into slice
			s = append(s, c)
		} else if len(s) == 0 {
			isBalanced = false
			break
		} else if bMap[s[len(s)-1]] != c {
			isBalanced = false
			break
		} else {
			//if the bracket matches deleting from slice
			s = s[:len(s)-1]
		}
	}

	if len(s) != 0 {
		isBalanced = false
	}

	return isBalanced
}
