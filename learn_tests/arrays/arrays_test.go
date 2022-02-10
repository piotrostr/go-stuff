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
	Describe("SumAll", func() {
		It("Should sum arrays and return array", func() {
			arr1 := []int{1, 2}
			arr2 := []int{0, 9}
			got := arrays.SumAll(arr1, arr2)
			want := []int{3, 9}
			Expect(got).To(Equal(want))
		})
	})
})
