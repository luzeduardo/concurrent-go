package main

import (
	"fmt"
	"sync"
)

type Semaphore struct {
	permits int
	cond    *sync.Cond
}

type WaitGrp struct {
	sema *Semaphore
}

func NewWaitGrp(size int) *WaitGrp {
	return &WaitGrp{sema: NewSemaphore(1 - size)}
}

func (wg *WaitGrp) Wait() {
	wg.sema.Acquire()
}

func (wg *WaitGrp) Done() {
	wg.sema.Release()
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
	wg := NewWaitGrp(4)
	for i := 0; i <= 4; i++ {
		go doWork(wg, i)
	}
	fmt.Println("All complete!")
}

func doWork(wg *WaitGrp, i int) {
	fmt.Println("Done working ", i)
	wg.Done()
}
