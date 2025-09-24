package main

import (
	"fmt"
	"sync"
)

/*
   sync.Map allows us to perform map operations concurrently
   without the problems of race conditions
   mehtods: load, store, delete, etc
*/

// SyncMap to demonstrate the function of sync.Map
func SyncMap() {
	//var m = make(map[string]string)
	var m sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		// Add key-val pairs to the map conncurrently
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			value := fmt.Sprintf("value%d", id)
			//m[key] = value
			m.Store(key, value)
		}(i)
	}
	wg.Wait()

	// Read values from the map concurrently
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		if val, ok := m.Load(key); ok {
			fmt.Printf("Key: %s, Value: %s\n", key, val)
		}
	}
}
