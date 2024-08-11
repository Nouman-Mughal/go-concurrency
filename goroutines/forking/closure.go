package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// closure1()
	closure2()
}

// goroutines run in the same address space, so the closure will be able to modify the variable.

func closure1() {
	salutation := "hello"

	wg.Add(1)

	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()

	wg.Wait()

	fmt.Println(salutation)
}

func closure2() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		// time.Sleep(3 * time.Second)
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()

	}
	wg.Wait()

}

/*
### What the Heap Does

The **heap** in computer science is a region of memory used for dynamically allocated objects whose size and lifetime
are not known at compile time. In contrast to the **stack**, which is used for static memory
allocation (like local variables in functions that are deallocated when the function returns), the heap is used
for objects that need to persist beyond the scope of the function that created them.

### Key Characteristics of the Heap

1. **Dynamic Memory Allocation**:
   - Objects are allocated on the heap when their size is not known at compile time, or when they need to live longer than the scope of a single function.
   - Memory on the heap must be manually managed, either by explicit deallocation (in languages like C/C++) or by garbage collection (in languages like Go, Java, etc.).

2. **Shared Memory**:
   - Since memory on the heap is globally accessible, multiple parts of a program (including different goroutines in Go) can share references to the same heap-allocated object.

3. **Lifetime**:
   - Objects on the heap can outlive the function or context that created them, depending on when they are explicitly freed or garbage collected.

### When and How References Get Saved to the Heap

In Go, variables can be allocated on the heap under certain conditions, such as when:

1. **Escape Analysis**:
   - Go uses a mechanism called **escape analysis** to decide whether a variable should be allocated on the stack or heap.
   - If a variable is needed after the function that created it returns (for example, if it’s referenced by a goroutine or returned from the function), Go’s compiler will allocate it on the heap.
   - Escape analysis determines this by checking if the variable "escapes" the function scope—meaning if it needs to persist beyond the function call.

2. **Closures and Goroutines**:
   - When a closure (like an anonymous function in a goroutine) captures a variable from its surrounding scope, and the closure outlives the function, the captured variable is moved to the heap.
   - This ensures that the variable remains accessible to the closure or goroutine even after the function that created it has exited.

### Example: Heap Allocation in Go

Consider the code example you provided:

```go
for _, salutation := range []string{"hello", "greetings", "good day"} {
    go func() {
        fmt.Println(salutation)
    }()
}
```

1. **Variable Capture**:
   - The `salutation` variable is captured by the anonymous function.
   - Because the goroutine may execute after the loop has completed, `salutation` cannot be stored on the stack. Otherwise, it would be deallocated when the loop iteration finishes.

2. **Heap Allocation**:
   - Go's escape analysis recognizes that `salutation` needs to persist beyond its local scope, so it allocates `salutation` on the heap.
   - The goroutines hold a reference to this heap-allocated `salutation` variable.

3. **Reference Management**:
   - As long as the goroutines are running and need access to `salutation`, the reference to it is kept alive, preventing garbage collection.
   - Once the goroutines complete and no more references to `salutation` exist, the Go garbage collector will eventually reclaim that heap memory.

### Summary

- The heap is used for dynamic memory allocation where the lifetime of variables is not bound to the scope of a function.
- In Go, variables are allocated on the heap when they "escape" their local function scope, such as when they are captured by a closure or need to persist after a function returns.
- The Go runtime handles this through escape analysis, ensuring variables are allocated on the heap when necessary and automatically manages memory through garbage collection.*/

func closure3() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		// time.Sleep(3 * time.Second)
		wg.Add(1)
		go func(salutation string) { // passing the copy of salutation to the goroutine so that goroutine can access the copy of salutation instead of the original salutation
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)

	}
	wg.Wait()

}
