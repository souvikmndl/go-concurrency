package main

import (
	"sync"
	"time"
)

// busy waiting --> (loop in consumer waitng for unlock)
// busy waiting increases cpu utilisation

// condvar.Broadcast() -> tells all goroutines at the same time
// condvar.Signal() -> signals one go routine at a time

var (
	mu      sync.Mutex // implements Locker iface
	condvar *sync.Cond
	data    int
)

// SyncCondition for sync.Cond
func SyncCondition() {
	condvar = sync.NewCond(&mu)

	go producer()
	go consumer()

	time.Sleep(5 * time.Second)
}

func producer() {
	for i := 0; i < 5; i++ {
		mu.Lock()
		data = i //produce
		// condvar.Signal()
		condvar.Broadcast()
		mu.Unlock()
		time.Sleep(time.Second)
	}
}

func consumer() {
	counter := 0
	for i := 0; i < 5; i++ {
		mu.Lock()
		for data != i {
			counter++ // checks cpu utilisation (how many iterations we had to wait for)
			condvar.Wait()
		}
		// without using singal/broadcast, this value would be vary high
		println(data, counter) //consume
		mu.Unlock()
	}
}
