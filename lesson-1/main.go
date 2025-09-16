package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
	}

	wg.Add(len(words)) // never hardcode this value

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg) // wg must always be passed as pointers as it should not be copied
	}

	wg.Wait()

	wg.Add(1)
	printSomething("Second print", &wg)
}
