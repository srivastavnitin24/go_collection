package mergesort

func Merge(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	//left := Merge(arr[:mid])
	//right := Merge(arr[mid:])
	//return sort(left, right)

	return sort(Merge(arr[:mid]), Merge(arr[mid:]))
}
func sort(left, right []int) []int {
	resArr := make([]int, len(left)+len(right))

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			resArr[k] = left[i]
			i++
		} else {
			resArr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		resArr[k] = left[i]
		k++
		i++
	}
	for j < len(right) {
		resArr[k] = right[j]
		j++
		k++
	}
	return resArr
}
