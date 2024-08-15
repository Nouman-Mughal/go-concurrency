// we are going to compare the performance of a goroutine with a thread.

// In my linux machine,
// taskset -c 0 perf bench sched pipe -T
// use this command to run the benchmark.
// cost of context switching is 1.65 microseconds.

// Now lets see the performance of goroutine context switching.

package main

import (
	"sync"
	"testing"
)

// func main() {
// 	BenchmarkContextSwitching()
// }

func BenchmarkContextSwitching(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})
	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}
	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}

// so when we run  ```GO111MODULE=off go test -bench=. -cpu=1 contextSwitching_test.go```

// I get 188.8 ns per context switch. which is 0.1888 microseconds.which means goroutines context switching is 8.7 times faster than thread context switching.
