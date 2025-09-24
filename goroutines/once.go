package main

import (
	"fmt"
	"sync"
)

var (
	once  sync.Once
	value int
)

func initialize() {
	value = 28
	fmt.Println("initialised value to ", value)
}

// SyncOnce demonstrates the use of sync.One
func SyncOnce() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(initialize)
			// even though this is inside a loop, initialize will be called only once
			fmt.Println(value)
			value++
		}()
		wg.Wait()
	}

}
