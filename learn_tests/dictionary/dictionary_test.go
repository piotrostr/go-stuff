package dictionary_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "learn_tests/dictionary"
)

var _ = Describe("Dictionary", func() {
	It("returns right word", func() {
		dict := Dict{"key": "value"}
		Expect(Search(dict, "key")).To(Equal("value"))
	})

	It("returns error if word is not there", func() {
		dict := Dict{"asf": "value"}
		value, err := Search(dict, "key")
		Expect(value).To(Equal(""))
		Expect(err).To(MatchError("no such word"))
	})
})
