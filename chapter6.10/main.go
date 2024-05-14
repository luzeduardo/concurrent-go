package main

import (
	"fmt"
	"sync"
	"time"
)

type Barrier struct {
	size      int
	waitCount int
	cond      sync.Cond
}

func NewBarrier(size int) *Barrier {
	condVar := sync.NewCond(&sync.Mutex{})
	return &Barrier{cond: *condVar}
}

func (b *Barrier) Wait() {
	b.cond.L.Lock()
	b.waitCount += 1

	if b.waitCount == b.size {
		b.waitCount = 0
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}

	b.cond.L.Unlock()
}

func workAndWait(name string, timeToWork int, barrier *Barrier) {
	for {
		fmt.Println(name, "is running")
		time.Sleep(time.Duration(timeToWork))
		fmt.Println(name, "is waiting on barrier")
		barrier.Wait()
	}
}

func main() {
	barrier := NewBarrier(2)

	go workAndWait("Red", 4, barrier)
	go workAndWait("Blue", 10, barrier)

	time.Sleep(time.Duration(60) * time.Second)
}
