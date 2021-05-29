package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	x := []int32{4, 2, 1, 3}
	minimumAbsDifference(x)

}

func minimumAbsDifference(numbers []int32) {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	size := len(numbers)
	minAbsDiff := math.MaxFloat64
	for i := 0; i < size-1; i++ {
		minAbsDiff = math.Min(minAbsDiff, math.Abs(float64(numbers[i]-numbers[i+1])))
	}

	k := 0
	for i := 0; i < size-1; i++ {
		if math.Abs(float64(numbers[i]-numbers[i+1])) == minAbsDiff {
			//str := strconv.Itoa(input[i]) + " " + strconv.Itoa(input[i+1])
			fmt.Println(strconv.Itoa(int(numbers[i])) + " " + strconv.Itoa(int(numbers[i+1])))
			//list = append(list, str, ":")
			k = k + 1
		}
	}

}
