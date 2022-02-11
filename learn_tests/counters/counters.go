package counters

import "sync"

type Counter struct {
	count int
	sync.Mutex
}

func (c *Counter) Count() int {
	return c.count
}

func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.count++
}
