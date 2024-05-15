package main

import (
	"fmt"
	"time"
)

func main() {
	msgChannel := make(chan string)
	go receiver(msgChannel)
	fmt.Println("Sending Hello...")

	msgChannel <- "Hello"

	fmt.Println("Sending there...")
	msgChannel <- "There"

	fmt.Println("Sending Stop...")
	msgChannel <- "Stop"
}

func receiver(messages chan string) {
	//when a goroutine push to a channel without another to read that messages
	//// this means that the sender will block until a receiver get ready to consume the message
	////and also channels are sync by default
	//
	//it results in a deadlock with the printed strings at the console below
	//  Sending Hello
	//  Receiver slept for 5 seconds
	//  fatal error: all goroutines are asleep - deadlock
	time.Sleep(5 * time.Second)
	fmt.Println("Receiver slept for 5 seconds")
}
