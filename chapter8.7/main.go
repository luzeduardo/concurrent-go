package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func sendMsgAfter(seconds time.Duration) <-chan string {
	messages := make(chan string)
	go func() {
		time.Sleep(seconds)
		messages <- "Hello"
	}()
	return messages
}

func main() {
	timeOutP, _ := strconv.Atoi(os.Args[1])
	messages := sendMsgAfter(10 * time.Second)

	timoutDuration := time.Duration(timeOutP) * time.Second
	select {
	case msg := <-messages:
		fmt.Println("Message received: ", msg)
	case tNow := <-time.After(timoutDuration):
		fmt.Println("Timed out. Waited until: ", tNow.Format("15:04:03"))
	}
}
