package main

import (
	"container/list"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// event listener that assigns each event to a shared data structure
func matchRecorder(matchEvents *list.List, mutex *sync.RWMutex) {
	for i := 0; ; i++ {
		mutex.Lock()
		matchEvents.PushBack("Match event " + strconv.Itoa(i))
		mutex.Unlock()

		time.Sleep(1 * time.Second)
		fmt.Println("Appended match event", strconv.Itoa(i))
	}
}

// simulates build of a response to send back to the user
func clientHandler(mEvents *list.List, mutex *sync.RWMutex, st time.Time) {
	mutex.RLock()
	allEvents := copyAllEvents(mEvents)
	mutex.RUnlock()

	timeTaken := time.Since(st)
	fmt.Println(len(allEvents), "events copied in", timeTaken)
}

func copyAllEvents(matchEvents *list.List) []string {
	i := 0
	allEvents := make([]string, matchEvents.Len())
	for e := matchEvents.Front(); e != nil; e = e.Next() {
		allEvents[i] = e.Value.(string)
		i++
	}
	return allEvents
}

func main() {
	mutex := sync.RWMutex{}
	matchEvents := list.New()
	// pre populate match events
	for j := 0; j < 10000; j++ {
		matchEvents.PushBack("Match Event")
	}

	go matchRecorder(matchEvents, &mutex)

	startTime := time.Now()
	for j := 0; j < 50000; j++ {
		go clientHandler(matchEvents, &mutex, startTime)
	}
	time.Sleep(100 * time.Second)
}
