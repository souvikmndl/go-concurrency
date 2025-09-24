package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// SyncAtomic demonstrates the use of sync/atomic to handle race conditions
func SyncAtomic() {
	var counter int32
	var wg sync.WaitGroup

	counter = 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//counter++ // wihtout atomic add, this will have race conditions
			atomic.AddInt32(&counter, 1) // this is much faster than sync.Mutex
		}()
	}
	wg.Wait()
	fmt.Println("Counter: ", counter)
	fmt.Println(atomic.LoadInt32(&counter))
}
