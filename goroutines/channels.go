package main

import "fmt"

// Channels has basics of channels
func Channels() {
	// var firstchan chan int
	firstchan := make(chan int)
	fmt.Println("Value of the channel: ", firstchan)
	fmt.Printf("Type of channel %T\n", firstchan)

	// reading and writing values to channels
	ch := make(chan int)
	fmt.Println("Hello from channels")
	go multiplyWithChannel(ch)
	ch <- 10 // passing/writing value to channel
	// this is also a blocking statement
	// this will wait until the value is read from the channel
	// if no one reads the value from this chan, execution will wait here
	// that is why if we pass the value before go func() in line 15, we would have reached a deadlock

	ch2 := make(chan int)
	close(ch2)            //close channel, no more read/writes
	elem, isOpen := <-ch2 // elem contains value inside ch, isOpen returns true if ch is open
	fmt.Println("IsOpen ", isOpen, " value ", elem)
	fmt.Println("Bye from main")

	c := make(chan string)
	go initStrings(c)
	for {
		resp, ok := <-c
		if ok == false {
			fmt.Println("Channel close ", ok)
			break
		}
		fmt.Println("Channel Open ", resp, ok)
	}

	testch := make(chan string)
	go func() {
		testch <- "value1"
		testch <- "value2"
		testch <- "value3"
		close(testch)
	}()

	// as long as there are values in channel, range will keep reading them
	// however is we dont close the channel in the ablove routine, it will be a deadlock
	// because range will try to read but there will be no values passed
	for res := range testch {
		fmt.Println(res)
	}

	// buffered channels: channels that can take n values at a time
	mychnl := make(chan string, 4)
	mychnl <- "abc"
	mychnl <- "cdef"
	mychnl <- "ghij"
	mychnl <- "xyz"
	fmt.Println("Length of ch ", len(mychnl), " capacity ", cap(mychnl))
	/*
	   - Buffered channels in Golang are a type of channel that can store a fixed number of values before a sender is blocked. Unlike unbuffered channels, which require a sender and receiver to be ready simultaneously for communication, buffered channels allow for a degree of decoupling between senders and receivers.

	   - A sender can transmit values to a buffered channel without waiting for an immediate receiver, as long as the buffer is not full. The send operation will only block if the channel's buffer has reached its capacity. So we can send upto capacity of channel and it wont block. After that sending values will block

	   - A receiver attempting to read from a buffered channel will block only if the channel's buffer is empty.

	   - Buffered channels facilitate asynchronous communication, allowing senders and receivers to operate more independently. This can be beneficial when you want to handle bursts of data or decouple the timing of different goroutines.
	*/
}

func multiplyWithChannel(ch chan int) {
	fmt.Println(100 * <-ch) // reading from channel
	// this is a blocking statement,
	// execution will wait here until a value is passed to ch
}

func initStrings(chnl chan string) {
	for v := 0; v < 3; v++ {
		chnl <- "go guru"
	}
	close(chnl)
}
