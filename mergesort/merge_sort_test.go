package mergesort

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
	It("Even no of array arguments", func() {
		sortedArray := Merge([]int{1, 8, 3, 0, 9, 5, 2})
		Expect(sortedArray).To(Equal([]int{0, 1, 2, 3, 5, 8, 9}))
	})
	It("Empty array", func() {
		sortedArray := Merge([]int{})
		Expect(sortedArray).To(Equal([]int{}))
	})
	It("Nil array", func() {
		var x []int
		sortedArray := Merge(nil)
		Expect(sortedArray).To(Equal(x))
	})
})
