package iteration_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"learn_tests/iteration"
	"testing"
)

var _ = Describe("Iteration", func() {
	It("repeats character a right", func() {
		Expect(iteration.Repeat("a")).To(Equal("aaaaa"))
	})

	It("repeats any character right", func() {
		Expect(iteration.Repeat("b")).To(Equal("bbbbb"))
		Expect(iteration.Repeat("c")).To(Equal("ccccc"))
	})
})

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a")
	}
}
