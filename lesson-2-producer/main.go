package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// NumberOfPizzas is the max number of pizzas producer will make
const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

// Producer represents the pizza producer
type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

// PizzaOrder represent details of each order
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			//pizza failed
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds\n", pizzaNumber, delay)
		//delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("** Ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("** The cook while cooking pizza #%d!", pizzaNumber)
		} else {
			success = true
			fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0

	// run forever or until we receive a quit notification

	// try to make pizzas
	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza (we sent something to the data channel)
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzeria is open for business")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer

	// print out the ending message
}
