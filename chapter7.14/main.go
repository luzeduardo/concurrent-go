package main

import (
	"container/list"
	"sync"
)

type Semaphore struct {
	permits int
	cond    sync.Cond
}

func NewSemaphore(permits int) *Semaphore {
	return &Semaphore{permits: permits, cond: *sync.NewCond(&sync.Mutex{})}
}

func (sema *Semaphore) Aquire() {
	sema.cond.L.Lock()

	sema.permits--
	if sema.permits <= 0 {
		sema.cond.Wait()
	}
	sema.cond.L.Unlock()
}

func (sema *Semaphore) Release() {
	sema.cond.L.Lock()
	sema.permits++
	sema.cond.Signal()
	sema.cond.L.Unlock()
}

type Channel[M any] struct {
	capacitySema *Semaphore
	sizeSema     *Semaphore
	mutex        sync.Mutex
	buffer       *list.List
}

func NewChannel[M any](capacity int) *Channel[M] {
	return &Channel[M]{
		capacitySema: NewSemaphore(capacity),
		sizeSema:     NewSemaphore(0),
		buffer:       list.New(),
	}
}

func (c *Channel[M]) Send(message M) {
	c.capacitySema.Aquire()

	c.mutex.Lock()
	c.buffer.PushBack(message)
	c.mutex.Unlock()

	c.sizeSema.Release()
}
