package main

import "fmt"

func main() {
	// assigning a nil to a channel blocks the channel from receiving or sending anything
	// it results in a deadlock
	var ch chan string = nil

	ch <- "message"

	fmt.Println("This is never printed")
}
