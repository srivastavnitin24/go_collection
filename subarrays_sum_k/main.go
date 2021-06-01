package main

import "fmt"

func main() {
	//findSubArray([]int{42, 15, 12, 8, 6, 32}, 26)
	//
	//findSubArray([]int{12, 5, 31, 13, 21, 8}, 49)
	//
	//findSubArray([]int{15, 51, 7, 81, 5, 11, 25}, 41)

	findSubArray([]int{23, 5, 1, 34, 12, 67, 9, 31, 6, 7, 27}, 40)

}

func findSubArray(inputArr []int, sum int) {

	outputArr := make([]int, len(inputArr))

	for i := 0; i < len(inputArr); i++ {
		temp := inputArr[i]
		outputArr = append(outputArr, temp)

		for j := i + 1; j < len(inputArr); j++ {

			if temp+inputArr[j] == sum {
				outputArr = append(outputArr, inputArr[j])
				fmt.Println(outputArr)
			} else if temp+inputArr[j] < sum {
				outputArr = append(outputArr, inputArr[j])
				temp = temp + inputArr[j]
			} else {
				outputArr = []int{}
				break
			}

		}
	}
}
