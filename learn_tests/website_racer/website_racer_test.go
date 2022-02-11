package website_racer_test

import (
	. "learn_tests/website_racer"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Racer", func() {
	It("races", func() {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		faster, err := Racer(slowServer.URL, fastServer.URL)
		Expect(err).To(BeNil())
		Expect(faster).To(Equal(fastServer.URL))
	})

	It("returns an error if the waiting time is over 10 seconds", func() {
		verySlowServer := makeDelayedServer(11 * time.Second)
		defer verySlowServer.Close()
		_, err := Racer(verySlowServer.URL, verySlowServer.URL)
		Expect(err).To(MatchError("timeout"))
	})
})

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(delay)
				w.WriteHeader(http.StatusOK)
			}))
}
