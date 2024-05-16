package main

import (
	"fmt"
	"time"
)

func receiver(msgs <-chan int) {
	// instead of check with a open flag reading from a channel, just use a range and it will iterate until the channel gets closed
	for msg := range msgs {
		fmt.Println(time.Now().Format("15:04:02"), "Received: ", msg)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Receiver finished")
}

func main() {
	msgChann := make(chan int)
	go receiver(msgChann)
	for i := 1; i < 4; i++ {
		fmt.Println(time.Now().Format("15:04:03"), "Sending: ", i)
		msgChann <- i
		time.Sleep(1 * time.Second)
	}
	close(msgChann)
	// after closing the channel the receiver keeps listening for messages in the channel
	// it results in values as zeroes (0) the default value for a closed channel of int.
	time.Sleep(5 * time.Second)
}
