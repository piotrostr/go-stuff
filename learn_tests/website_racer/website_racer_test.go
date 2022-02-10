package website_racer_test

import (
	"learn_tests/website_racer"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("WebsiteRacer", func() {
	It("races", func() {
		slowURL := "http://www.facebook.com"
		fastURL := "http://www.quii.dev"

		got := website_racer.Racer(slowURL, fastURL)

		Expect(got).To(Equal(fastURL))
	})
})
