package integers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"learn_tests/integers"
)

var _ = Describe("Integers", func() {
	It("adds", func() {
		Expect(integers.Add(2, 2)).To(Equal(4))
		Expect(integers.Add(3, 2)).To(Equal(5))
	})
})
