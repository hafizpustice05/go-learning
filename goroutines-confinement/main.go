package main

import (
	"fmt"
	"sync"
	"time"
)

/*
A function that is responsible for managing ticket allocations and responds to user request
It listen for incoming request on a chanel(ticketChan) and signals
on another chanel (doneChan) when it's time to stop
*/
func manageTicket(ticketChan chan int, doneChan chan bool, tickets *int) {
	for {
		select {
		case userId := <-ticketChan:
			if *tickets > 0 {
				*tickets--
				fmt.Printf("User %d purchased the ticket and remainging ticket %d.\n", userId, *tickets)
			} else {
				fmt.Printf("User %d Found no ticket\n", userId)
			}
		case <-doneChan:
			fmt.Printf("Tickets remaining: %d\n", *tickets)
		}
	}
}

/*
A function that simulates a user trying to buy a ticket.
It sends a request to the manageTicket goroutine through ticketChanel
*/

func buyTicket(wg *sync.WaitGroup, ticketChan chan int, userId int) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	ticketChan <- userId
}

func main() {

	now := time.Now()

	var wg sync.WaitGroup        //waitgroup to wait for all goroutines to finish
	tickets := 5                 //number of ticket
	ticketChan := make(chan int) //ticket chanel for sending ticket purchase request
	doneChan := make(chan bool)  //Chanel for signaling to stop

	go manageTicket(ticketChan, doneChan, &tickets)

	for userId := 0; userId <= 2000; userId++ {
		wg.Add(1)
		go buyTicket(&wg, ticketChan, userId)
	}

	wg.Wait()

	doneChan <- true

	fmt.Printf("execution time: ", time.Since(now))
}
