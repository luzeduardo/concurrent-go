package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int, mutex *sync.Mutex) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			mutex.Lock()
			frequency[cIndex] += 1
			mutex.Unlock()
		}
	}
	fmt.Println("Completed: ", url)
}

func main() {
	//all goroutines will share the same data structure in memory
	//and it will add a race condition
	var frequency = make([]int, 26)
	var mutex = sync.Mutex{}
	for i := 1000; i <= 1020; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency, &mutex)
	}
	//forcing the main thread to not finish before goroutines finishes their jobs
	time.Sleep(5 * time.Second)
	var allLetterSum int
	for i, c := range allLetters {
		mutex.Lock()
		allLetterSum += frequency[i]
		fmt.Printf("%c-%d\n", c, frequency[i])
		mutex.Unlock()
	}
	fmt.Printf("Letters in docs: %d\n", allLetterSum)
}
