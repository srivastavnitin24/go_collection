package binarysearch

func nextLetterInAlphabet(arr []string, find string) string {
	if len(arr) == 0 {
		return ""
	}
	res := ""
	start, end := 0, len(arr)-1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == find {
			res = arr[mid]
			return res
		}
		if arr[mid] < find {
			start = mid + 1
		}
		if arr[mid] > find {
			res = arr[mid]
			end = mid - 1
		}
	}
	return res
}
