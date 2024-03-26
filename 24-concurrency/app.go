package main

import (
	"fmt"
	"time"
)

// parellelization using go routines (non-blocking calls)
func main() {
	doneChans := make([]chan bool, 4)
	for i := range doneChans {
		doneChans[i] = make(chan bool)
	}
	go slowGreet("How ... are ... you ...?", doneChans[0])
	go greet("Nice to meet you!", doneChans[1])
	go greet("How are you?", doneChans[2])
	go greet("I hope you're liking the course!", doneChans[3])

	for _, doneChan := range doneChans {
		<-doneChan
	}

	// doneChan := make(chan bool)
	// go slowGreet("How ... are ... you ...?", doneChan)
	// go greet("Nice to meet you!", doneChan)
	// go greet("How are you?", doneChan)
	// go greet("I hope you're liking the course!", doneChan)
	// <-doneChan
	// <-doneChan
	// <-doneChan
	// <-doneChan
}

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
	done <- true
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	done <- true
	// close(done) // Can be called when the channel is no longer needed
}
