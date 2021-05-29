package binarysearch

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Suite")
}

var _ = Describe("Merge sort", func() {
	//It("Even no of array arguments", func() {
	//	sortedArray := binarySearch([]int{2, 4, 6, 8, 10, 11, 13}, 4)
	//	Expect(sortedArray).To(Equal(true))
	//})
	//It("array with 1 element", func() {
	//	sortedArray := binarySearch([]int{5}, 5)
	//	Expect(sortedArray).To(Equal(true))
	//})
	//It("array with 1 element", func() {
	//	sortedArray := binarySearch([]int{5}, 2)
	//	Expect(sortedArray).To(Equal(false))
	//})
	//It("Empty array", func() {
	//	sortedArray := binarySearch([]int{}, 1)
	//	Expect(sortedArray).To(Equal(false))
	//})
	//It("Nil array", func() {
	//	var x []int
	//	sortedArray := binarySearch(x, 0)
	//	Expect(sortedArray).To(Equal(false))
	//})
	//It("Decreasing sorted array", func() {
	//	sortedArray := binarySearch([]int{13, 11, 10, 8, 6, 4, 2}, 13)
	//	Expect(sortedArray).To(Equal(true))
	//})
	//It("Find first occurrence in sorted array", func() {
	//	sortedArray := binarySearchFristOccurrence([]int{2, 4, 4, 6, 8, 10, 11, 13}, 4)
	//	Expect(sortedArray).To(Equal(1))
	//})
	//It("Find first occurrence in sorted array", func() {
	//	sortedArray := binarySearchFristOccurrence([]int{2, 4, 4, 6, 8, 10, 11, 13}, 5)
	//	Expect(sortedArray).To(Equal(-1))
	//})
	//It("Find first occurrence in sorted array", func() {
	//	sortedArray := binarySearchFristOccurrence([]int{2, 4, 4, 6, 8, 10, 11, 13, 17, 17}, 17)
	//	Expect(sortedArray).To(Equal(8))
	//})
	//It("Find first occurrence in sorted array", func() {
	//	sortedArray := binarySearchFristOccurrence([]int{}, 10)
	//	Expect(sortedArray).To(Equal(-1))
	//})
	//It("Find first occurrence in sorted array", func() {
	//	sortedArray := find_index_in_rotated_array([]int{11, 12, 15, 1, 2, 5, 6, 8})
	//	Expect(sortedArray).To(Equal(3))
	//})
	//It("Find first occurrence in sorted array", func() {
	//	sortedArray := find_index_in_rotated_array([]int{15, 1, 2, 5, 6, 8, 11, 13})
	//	Expect(sortedArray).To(Equal(1))
	//})
	//It("Find first occurrence in sorted array", func() {
	//	sortedArray := find_index_in_rotated_array([]int{2, 3, 5, 6, 8, 11, 12, 15, 1})
	//	Expect(sortedArray).To(Equal(8))
	//})
	//It("Find first occurrence in sorted array", func() {
	//	sortedArray := find_index_in_rotated_array([]int{2})
	//	Expect(sortedArray).To(Equal(-1))
	//})
	//It("Find first occurrence in sorted array", func() {
	//	sortedArray := find_index_in_rotated_array([]int{})
	//	Expect(sortedArray).To(Equal(-1))
	//})
	//It("Find first occurrence in sorted array", func() {
	//	sortedArray := find_index_in_rotated_array([]int{2, 3, 5, 6, 8, 11, 12, 15, 1})
	//	Expect(sortedArray).To(Equal(8))
	//})
	//
	//It("findElementInRotatedSortedArray", func() {
	//	sortedArray := findElementInRotatedSortedArray([]int{2, 3, 5, 6, 8, 11, 12, 15, 1}, 3)
	//	Expect(sortedArray).To(Equal(1))
	//})
	//It("findElementInRotatedSortedArray", func() {
	//	sortedArray := findElementInRotatedSortedArray([]int{2, 3, 5, 6, 8, 11, 12, 15, 1}, 4)
	//	Expect(sortedArray).To(Equal(-1))
	//})
	//
	//It("searchingInNearlySortedArray", func() {
	//	sortedArray := searchingInNearlySortedArray([]int{5, 10, 30, 20, 40}, 10)
	//	Expect(sortedArray).To(Equal(1))
	//})
	//It("searchingInNearlySortedArray", func() {
	//	sortedArray := searchingInNearlySortedArray([]int{5, 10, 20, 30, 40}, 40)
	//	Expect(sortedArray).To(Equal(4))
	//})
	//
	//It("floorofelementinsortedarray", func() {
	//	sortedArray := floorofelementinsortedarray([]int{5, 10, 20, 30, 40}, 12)
	//	Expect(sortedArray).To(Equal(10))
	//})
	//It("floorofelementinsortedarray", func() {
	//	sortedArray := floorofelementinsortedarray([]int{5, 10, 20, 30, 40}, 42)
	//	Expect(sortedArray).To(Equal(40))
	//})
	//It("floorofelementinsortedarray", func() {
	//	sortedArray := floorofelementinsortedarray([]int{1}, 0)
	//	Expect(sortedArray).To(Equal(-1))
	//})

	//It("ceilofelementinsortedarray", func() {
	//	sortedArray := ceilofelementinsortedarray([]int{5, 10, 20, 30, 40}, 12)
	//	Expect(sortedArray).To(Equal(20))
	//})
	//It("ceilofelementinsortedarray", func() {
	//	sortedArray := ceilofelementinsortedarray([]int{1, 2, 3, 4, 8, 10, 10, 12, 19}, 5)
	//	Expect(sortedArray).To(Equal(8))
	//})
	It("nextLetterInAlphabet", func() {
		sortedArray := nextLetterInAlphabet([]string{"a", "c", "e", "g", "h"}, "h")
		Expect(sortedArray).To(Equal("h"))
	})
	//It("nextLetterInAlphabet", func() {
	//	sortedArray := nextLetterInAlphabet([]int{1, 2, 3, 4, 8, 10, 10, 12, 19}, 5)
	//	Expect(sortedArray).To(Equal(8))
	//})

})
