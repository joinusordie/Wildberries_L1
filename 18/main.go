package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Counter struct {
	sync.Mutex
	val int
}

func NewCounter() *Counter {
	return &Counter{
		val: 0,
	}
}

func (c *Counter) inc() {
	c.Lock()
	c.val++
	c.Unlock()
}

func main() {
	var wg sync.WaitGroup

	counter := NewCounter()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.inc()
		}()
	}

	wg.Wait()
	fmt.Println(counter.val)
}
