package chapter67

import (
	"sync"
)

type WGroup struct {
	groupSize int
	cond      *sync.Cond
}

func NewWaitGroup() *WGroup {
	return &WGroup{cond: sync.NewCond(&sync.Mutex{})}
}

func (wg *WGroup) Add(delta int) {
	wg.cond.L.Lock()
	wg.groupSize += delta
	wg.cond.L.Unlock()
}

func (wg *WGroup) Wait() {
	wg.cond.L.Lock()
	for wg.groupSize > 0 {
		wg.cond.Wait()
	}
	wg.cond.L.Unlock()
}

func (wg *WGroup) Done() {
	wg.cond.L.Lock()

	wg.groupSize--

	if wg.groupSize == 0 {
		wg.cond.Broadcast()
	}
	wg.cond.L.Unlock()
}
