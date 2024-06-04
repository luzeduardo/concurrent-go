package main

import "fmt"

func printNumbers(numbers <-chan int, quit chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-numbers)
		}
		// pattern quitting channel
		close(quit)
	}()
}

func main() {
	numbers := make(chan int)
	quit := make(chan int)
	printNumbers(numbers, quit)

	next := 10
	for i := 1; ; i++ {
		next += 1
		select {
		// passing copies of numbers to the channel, not sharing memory
		case numbers <- next:
			fmt.Println("In case")
		case <-quit:
			fmt.Println("Quitting number generation")
			return
		}
	}
}
