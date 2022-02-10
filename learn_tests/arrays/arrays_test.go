package arrays_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"learn_tests/arrays"
)

var _ = Describe("Arrays", func() {
	Describe("Sum", func() {
		It("Should sum right for nums equal 15", func() {
			numbers := []int{1, 2, 3, 4, 5}
			Expect(arrays.Sum(numbers)).To(Equal(15))
		})

		It("Should sum right for nums equal any sum", func() {
			numbers := []int{2, 2, 3, 4, 5}
			Expect(arrays.Sum(numbers)).To(Equal(16))
		})

		It("Should sum right for any array size", func() {
			numbers := []int{1, 1, 1, 1, 1, 1, 1}
			Expect(arrays.Sum(numbers)).To(Equal(7))
		})
	})
})
