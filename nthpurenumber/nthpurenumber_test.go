package nthpurenumber

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
		sortedArray := nthPureNumber(5)
		Expect(sortedArray).To(Equal(5445))
	})
})
