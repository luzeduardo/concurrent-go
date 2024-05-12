package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// the sequential-to-parallel ration will limit the performance scalability
	// it is essential to reduce the time spent holding the mutex lock
	mutex.Lock()
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	mutex.Unlock()
	fmt.Println("Completed: ", url)
	wg.Done()
}

func main() {
	// all goroutines will share the same data structure in memory
	// and it will add a race condition
	frequency := make([]int, 26)
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(21)
	for i := 1000; i <= 1020; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency, &mutex, &wg)
	}
	// forcing the main thread to not finish before goroutines finishes their jobs
	var allLetterSum int
	for i, c := range allLetters {
		wg.Wait()
		mutex.Lock()
		allLetterSum += frequency[i]
		fmt.Printf("%c-%d\n", c, frequency[i])
		mutex.Unlock()
	}
	fmt.Printf("Letters in docs: %d\n", allLetterSum)
}
