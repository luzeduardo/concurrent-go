package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	fmt.Println("Completed: ", url)
}

func main() {
	//all goroutines will share the same data structure in memory
	//and it will add a race condition
	var frequency = make([]int, 26)
	for i := 1000; i <= 1020; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency)
	}
	//forcing the main thread to not finish before goroutines finishes their jobs
	time.Sleep(5 * time.Second)

	for i, c := range allLetters {
		fmt.Printf("%c-%d\n", c, frequency[i])
	}
}
