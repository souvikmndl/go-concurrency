package main

import (
	"fmt"
	"sync"
	"time"
)

// Waitgroups can be used to control Channels, along with

var wg sync.WaitGroup

func worker(id int, check chan bool) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	check <- true // signals that the routine has completed its work
}

// Waitgroups to control channels
func Waitgroups() {
	chan1 := make(chan bool)
	for i := 1; i <= 5; i++ {
		//wg.Add(1)
		go func(chan bool) {
			//defer wg.Done()
			worker(i, chan1)
		}(chan1)
		<-chan1 // works like wg.Wait()
		//thread will wait here till a value is passed to the chan
		//that it can read from
	}
	//wg.Wait()
}
