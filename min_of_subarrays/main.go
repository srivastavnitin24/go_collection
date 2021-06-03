//{4,5,6,2,-1,5,7}
//Sublength 3
//4 2 -1
package main

import "fmt"

func main() {
	arr := []int{4, 5, 6, 2, -1, 5, 7}
	sub := 3
	x := findMin(arr, sub)
	fmt.Println(x)
}

func findMin(arr []int, sub int) []int {
	var tempArr []int
	for i := 0; i <= len(arr)-sub; i++ {
		temp := arr[i] //4
		for j := i; j < sub+i; j++ {
			if temp > arr[j] {
				temp = arr[j]
			}

		}
		tempArr = append(tempArr, temp)
		temp = 0
	}
	return tempArr
}
