package main

import (
	"fmt"
	"sync"
	"time"
)

func stingy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*money += 10
		mutex.Unlock()
	}
	fmt.Println("Stingy done")
}

func spendy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*money -= 10
		mutex.Unlock()
	}
	fmt.Println("Speny done")
}

func main() {
	money := 100
	mutex := sync.Mutex{}
	go stingy(&money, &mutex)
	go spendy(&money, &mutex)
	time.Sleep(2 * time.Second)
	mutex.Lock()
	println("Money in banck account: ", money)
	mutex.Unlock()
}
