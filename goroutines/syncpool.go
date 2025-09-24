package main

import (
	"fmt"
	"sync"
)

type SomeObject struct {
	Data []byte
}

func createObject() *SomeObject {
	return &SomeObject{
		Data: make([]byte, 1024*1024), // 1 MB of data
	}
}

// SyncPool demonstrates the basics of sync.Pool
func SyncPool() {
	var memoryPiece int //0
	// allocates memory and clears it (after use) very efficiently
	objectPool := sync.Pool{
		New: func() interface{} {
			memoryPiece++
			return createObject()
		},
	}

	const worker = 1024 * 1024
	var wg sync.WaitGroup
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			obj := objectPool.Get().(*SomeObject)
			// some use to do stuff
			objectPool.Put(obj) // adds to the pool
		}()
	}
	wg.Wait()

	// the number of workers is huge. But sync.Pool creates and reuses memory efficiently
	// memoryPiece++ counts the number of memories created, and it will be far less than 1024*1024
	// thats the magic of sync.Pool --> you create memory inside pool using New: func() interface{}
	// and after use you put it back to the pool using objectPool.Put(obj) so that the pool can reuse it
	fmt.Println("Done ", memoryPiece)
}
