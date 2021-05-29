package gorouteschannels

import "testing"

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Suite")
}

var _ = Describe("Merge sort", func() {
	It("Even no of array arguments", func() {
		n := mainfunc()
		Expect(n).To(Equal(55))
	})
})
