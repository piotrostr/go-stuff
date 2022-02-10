package website_racer_test

import (
	"learn_tests/website_racer"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("WebsiteRacer", func() {
	It("races", func() {
		slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))

		fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		got := website_racer.Racer(slowURL, fastURL)

		Expect(got).To(Equal(fastURL))

		slowServer.Close()
		fastServer.Close()
	})
})
