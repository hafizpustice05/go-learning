package main

import (
	"fmt"
	"sync"
)

// Mutex (Mutual Exclusion)
// A mutex (mutual exclusion lock) is used to protect shared resources
// from concurrent access, preventing race conditions.
var mutex sync.Mutex

func buyTickets(wg *sync.WaitGroup, userId int, remainingTickets *int) {
	defer wg.Done()
	mutex.Lock()
	if *remainingTickets > 0 {
		*remainingTickets-- //user purchases a ticket
		fmt.Printf("user %d purchased a ticket. Tickets remaining: %d\n", userId, *remainingTickets)
	} else {
		fmt.Printf("User %d Found no ticket\n", userId)
	}
	mutex.Unlock()
}
func main() {
	fmt.Println("hello")

	var tickets int = 500
	var wg sync.WaitGroup

	//Simulating a lot of users trying to buy a tickets
	// For simlicity, let's assume number of tickerts a user can buy a one ticket
	for userId := 0; userId <= 2000; userId++ {
		wg.Add(1)
		go buyTickets(&wg, userId, &tickets)
	}

	wg.Wait()
}
