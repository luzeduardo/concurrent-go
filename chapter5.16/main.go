package main

import (
	"fmt"
	"sync"
)

type Semaphore struct {
	permits int
	cond    *sync.Cond
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		permits: n,
		cond:    sync.NewCond(&sync.Mutex{}),
	}
}

func (rw *Semaphore) Acquire() {
	rw.cond.L.Lock()

	for rw.permits <= 0 {
		rw.cond.Wait()
	}

	rw.permits--
	rw.cond.L.Unlock()
}

func (rw *Semaphore) Release() {
	rw.cond.L.Lock()

	rw.permits++
	rw.cond.Signal()

	rw.cond.L.Unlock()
}

func main() {
	semaphore := NewSemaphore(0)
	for i := 0; i < 20; i++ {
		go doWork(semaphore, i)

		semaphore.Acquire()
		fmt.Println("Child goroutine finished: ", i)
	}
}

func doWork(semaphore *Semaphore, i int) {
	fmt.Println("Work started: ", i)
	fmt.Println("Work Finished ", i)

	semaphore.Release()
}
