package main

import (
	"fmt"
	"sync"
	"time"
)

func receiver(messages <-chan int, wg *sync.WaitGroup) {
	for {
		msg := <-messages
		fmt.Println(time.Now().Format("13:01:03"), "Received: ", msg)
		if msg == 5 {
			wg.Done()
		}
	}
}

func sender(messages chan<- int, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		fmt.Println(time.Now().Format("15:03:05"), "Sending: ", i)
		messages <- i
		time.Sleep(1 * time.Second)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	chnT := make(chan int)
	go sender(chnT, &wg)
	go receiver(chnT, &wg)
	wg.Wait()
}
