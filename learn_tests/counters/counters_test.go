package counters_test

import (
	"learn_tests/counters"
	"sync"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sync", func() {
	var counter *counters.Counter

	BeforeEach(func() {
		counter = &counters.Counter{}
		Expect(counter).NotTo(BeNil())
		Expect(counter.Count()).To(Equal(0))
	})

	It("counts right", func() {
		counter.Increment()
		Expect(counter.Count()).To(Equal(1))
		counter.Increment()
		Expect(counter.Count()).To(Equal(2))
	})

	It("runs safe concurrently", func() {
		wantedCount := 1000
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Increment()
				wg.Done()
			}()
		}

		wg.Wait()
		Expect(counter.Count()).To(Equal(wantedCount))
	})
})
