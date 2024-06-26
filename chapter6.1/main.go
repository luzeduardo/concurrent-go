package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i <= 10; i++ {
		go doWork(i, &wg)
	}

	wg.Wait()
	fmt.Println("All complete!")
}

func doWork(id int, wg *sync.WaitGroup) {
	i := rand.Intn(5)

	time.Sleep(time.Duration(i) * time.Second)

	fmt.Println(id, "done working after ", i, " seconds")
	wg.Done()
}
