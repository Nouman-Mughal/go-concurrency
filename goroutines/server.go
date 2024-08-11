package main

import "sync"

func main() {
	// WrongFork()
	RightFork()
}

func WrongFork() {

	sayHello := func() {
		println("Hello")
	}
	go sayHello()
	// The go routine will be created and scheduled with go runtime to execute but it may not get the chance to run before the
	// main function returns.
}

func RightFork() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		println("Hello")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()
}
