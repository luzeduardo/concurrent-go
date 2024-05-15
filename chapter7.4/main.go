package main

import (
	"fmt"
	"time"
)

func main() {
	msgChannel := make(chan string)
	go sender(msgChannel)

	fmt.Println("Reading message from channel...")

	msg := <-msgChannel
	fmt.Print("Received: ", msg)
}

func sender(message chan string) {
	time.Sleep(5 * time.Second)
	fmt.Println("Sender slept for 5 seconds")
}
