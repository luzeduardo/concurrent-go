package main

import (
	"fmt"
	"sync"
	"time"
)

func receiver(messages chan int, wGroup *sync.WaitGroup) {
	msg := 0
	for msg != -1 {
		time.Sleep(1 * time.Second)
		msg = <-messages
		fmt.Println("Received: ", msg)
	}
	wGroup.Done()
}

func main() {
	msgChannel := make(chan int, 3)
	wgGroup := sync.WaitGroup{}
	wgGroup.Add(1)

	go receiver(msgChannel, &wgGroup)
	// it will keep adding elements to the channel until it reaches the limit of the buffer,
	// so it will be blocked waiting for the receiver to start receiving those messages, emptying the buffer and open space for
	// new messages to be emitted

	for i := 1; i <= 6; i++ {
		fmt.Println(time.Now().Format("15:04:05"), "Sending:", i)
		msgChannel <- i
	}

	msgChannel <- -1
	wgGroup.Wait()
}
