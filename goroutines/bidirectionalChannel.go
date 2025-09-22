package main

import "fmt"

// BiDirectionalChannels
func BiDirectionalChannels() {
	// the channels we created in channels.go are bidirectional
	// we can read froma nd write to them

	// chanl1 := make(<-chan string) // receiving channel, only to read values from
	// chanl2 := make(chan<- string) // sending channel, can. only send values to

	chan1 := make(chan string)
	chan2 := make(chan string)

	go sending(chan1)
	valueFromChannel := <-chan1
	fmt.Println("value from channel ", valueFromChannel)
	go receiving(chan2)
	chan2 <- valueFromChannel

	chanl1 := make(chan string) // bidirectional
	go convert(chanl1)          // we can pass a bidirectional channel to a func that only takes
	// unidirectional channel. Inside the func this will be treated as unidirectional only
	fmt.Println(<-chanl1)
}

func sending(s chan string) {
	s <- "Go Guru"
}

func receiving(c chan string) {
	fmt.Println(<-c)
}

// send only channel
func convert(s chan<- string) {
	s <- "some value"
}
