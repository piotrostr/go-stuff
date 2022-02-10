package shapes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"learn_tests/shapes"
)

var _ = Describe("Shapes", func() {
	Describe("Perimeter", func() {
		It("should calculate perimeter right", func() {
			rectangle := shapes.Rectangle{10.0, 10.0}
			got := shapes.Perimeter(rectangle)
			want := 40.0
			Expect(got).To(Equal(want))
		})
	})

	Describe("Area", func() {
		It("should calculate area right", func() {
			rectangle := shapes.Rectangle{5.0, 10.0}
			got := shapes.Area(rectangle)
			want := 50.0
			Expect(got).To(Equal(want))
		})
	})
})
