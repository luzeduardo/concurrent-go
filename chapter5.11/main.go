package main

import (
	"fmt"
	"sync"
	"time"
)

type ReadWriterMutex struct {
	readersCounter int
	writersWaiting int
	writerActive   bool
	cond           *sync.Cond
}

func NewReadWriteMutex() *ReadWriterMutex {
	return &ReadWriterMutex{cond: sync.NewCond(&sync.Mutex{})}
}

func (rw *ReadWriterMutex) ReadLock() {
	rw.cond.L.Lock()

	for rw.writersWaiting > 0 || rw.writerActive {
		rw.cond.Wait()
	}

	rw.readersCounter++
	rw.cond.L.Unlock()
}

func (rw *ReadWriterMutex) WriteLock() {
	rw.cond.L.Lock()

	rw.writersWaiting++

	for rw.readersCounter > 0 || rw.writerActive {
		rw.cond.Wait()
	}

	rw.writersWaiting--
	rw.writerActive = true
	rw.cond.L.Unlock()
}

func (rw *ReadWriterMutex) ReadUnlock() {
	rw.cond.L.Lock()
	rw.readersCounter--

	if rw.readersCounter == 0 {
		rw.cond.Broadcast()
	}

	rw.cond.L.Unlock()
}

func (rw *ReadWriterMutex) WriterUnlock() {
	rw.cond.L.Lock()
	rw.writerActive = false
	rw.cond.Broadcast()
	rw.cond.L.Unlock()
}

func main() {
	rwMutex := ReadWriterMutex{}

	for i := 0; i < 2; i++ {
		go func() {
			for {
				rwMutex.ReadLock()
				time.Sleep(1 * time.Second)
				fmt.Println("Read done")
				rwMutex.ReadUnlock()
			}
		}()
	}
	time.Sleep(1 * time.Second)
	rwMutex.WriteLock()
	fmt.Println("Write finished")
}
