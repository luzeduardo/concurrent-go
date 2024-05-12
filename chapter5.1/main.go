package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func stingy(money *int, cond *sync.Cond) {
	for i := 0; i < 1000000; i++ {
		cond.L.Lock()
		*money += 10
		// Signal -> After updating the shared state above,  the condition unlocks the mutex
		//   - So other gouroutine will receive this Signal and wake up, and reaqcquire the mutex to proceed with its job
		if *money >= 50 {
			cond.Signal()
		}
		cond.L.Unlock()
	}
	fmt.Println("Stingy Done")
}

func spendy(money *int, cond *sync.Cond) {
	for i := 0; i < 200000; i++ {
		cond.L.Lock()

		for *money < 50 {
			// Wait -> performs two operations atomically:
			//  - releases the mutex
			//  - blocks the current execution, effectivelly putting the goroutine to sleep
			//  Since the mutex is now available the other goroutine can acquire it to continue with its execution
			//
			//  So the wait function releases the mutex and suspends the current goroutine execution in an atomic manner
			cond.Wait()
		}

		*money -= 50
		if *money < 0 {
			fmt.Println("Money is negative!")
			os.Exit(1)
		}
		cond.L.Unlock()
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)

	go stingy(&money, cond)
	go spendy(&money, cond)
	time.Sleep(2 * time.Second)
	mutex.Lock()
	println("Money in bank account: ", money)
	mutex.Unlock()
}
