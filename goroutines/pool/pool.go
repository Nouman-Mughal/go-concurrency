package main

import (
	"fmt"
	"sync"
)

func main() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024) // 1kb
			return &mem
		},
	}
	// seeding the pool by 4kb memory.
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024

	var wg sync.WaitGroup

	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
			// Assume here we done something with the memory quickly
		}()
	}

	wg.Wait()

	fmt.Printf("%d calculators were created", numCalcsCreated)

}

// only 13 calculators were created on my machine.

// for 1048576 number of goroutines we could consume GBs of memory. Pool implementations saved all the cost of memory
// and we reuse the same memory all this time.
