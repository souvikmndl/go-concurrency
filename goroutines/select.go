package main

import (
	"fmt"
)

// SelectChan explores select statements
func SelectChan() {
	// the cases refer to communications, send or receive
	chanl1 := make(chan string)
	chanl2 := make(chan int)

	go funcTwo(chanl2)
	go funcOne(chanl1)

	// as soon as a chanl is sent a value, it is received here
	// select will read the first value it receives and ignores the other one
	// after the first value is received, select's work is done and exec moves
	// on to the next block of code
	select {
	case val1 := <-chanl1:
		fmt.Println(val1)
	case val2 := <-chanl2:
		fmt.Println(val2)
		// default:
		// 	fmt.Println("Some error")
	}
}

func funcOne(chanl1 chan string) {
	// time.Sleep(1 * time.Second)
	chanl1 <- "Welcome to channel 1/funcOne"
}

func funcTwo(chanl2 chan int) {
	// time.Sleep(1 * time.Second)
	chanl2 <- 20
}
