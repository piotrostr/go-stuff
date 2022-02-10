package shapes_test

import (
	"learn_tests/shapes"
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shapes", func() {
	Describe("Perimeter", func() {
		It("should calculate perimeter right", func() {
			rectangle := shapes.Rectangle{10.0, 10.0}
			got := rectangle.Perimeter()
			want := 40.0
			Expect(got).To(Equal(want))
		})
	})

	Describe("Area", func() {
		It("should calculate rectangle area right", func() {
			rectangle := shapes.Rectangle{5.0, 10.0}
			got := rectangle.Area()
			want := 50.0
			Expect(got).To(Equal(want))
		})

		It("should calculate circle area right", func() {
			circle := shapes.Circle{10.0}
			got := circle.Area()
			want := 100.0 * math.Pi
			Expect(got).To(Equal(want))
		})

		It("should calculate triangle shape right", func() {
			triangle := shapes.Triangle{10.0, 10.0}
			got := triangle.Area()
			want := 10.0 * 10.0 * 0.5
			Expect(got).To(Equal(want))
		})

		Describe("parametrized cases", func() {
			areaTests := []struct {
				name  string
				shape shapes.Shape
				want  float64
			}{
				{
					"Rectangle",
					shapes.Rectangle{12, 6},
					72.0,
				},
				{
					"Circle",
					shapes.Circle{10},
					314.15926535897932,
				},
				{
					"Triangle",
					shapes.Circle{100},
					31415.926535897932,
				},
			}

			for _, tc := range areaTests {
				It("parametrized case for "+tc.name, func() {
					got := tc.shape.Area()
					Expect(got).To(Equal(tc.want))
				})
			}
		})
	})
})
