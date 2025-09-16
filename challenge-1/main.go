package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup
	msg = "Hello, world!"

	wg.Add(1)
	updateMessage("Hello universe", &wg)
	printMessage()
	wg.Wait()

	wg.Add(1)
	updateMessage("Hello cosmos", &wg)
	printMessage()
	wg.Wait()

	wg.Add(1)
	updateMessage("Hello world", &wg)
	printMessage()
	wg.Wait()
}
