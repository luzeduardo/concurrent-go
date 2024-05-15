package main

import "fmt"

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
	msg := ""
	for msg != "Stop" {
		msg = <-messages
		fmt.Println("Received: ", msg)
	}
}
