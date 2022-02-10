package concurrency_test

import (
	"learn_tests/concurrency"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

var _ = Describe("Concurrency", func() {
	It("checks websites right", func() {
		websites := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}
		got := concurrency.CheckWebsites(mockWebsiteChecker, websites)
		want := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}
		Expect(got).To(Equal(want))
	})
})

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrency.CheckWebsites(slowWebsiteChecker, urls)
	}
}
