package main

import "fmt"

// Count a number of vowels and consonants in a given String?
func main() {
	str := "india"
	countVol := vowel(str)
	countCon := len(str) - countVol
	fmt.Printf("count of vowel %v and count of consonant : %v", countVol, countCon)

}

func vowel(str string) int {
	count := 0
	for _, i2 := range str {
		if i2 == 'a' || i2 == 'e' || i2 == 'i' || i2 == 'o' || i2 == 'u' {
			count++
		}
	}
	return count
}
