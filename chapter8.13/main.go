package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string) <-chan []int {
	result := make(chan []int)
	go func() {
		defer close(result)
		frequency := make([]int, 26)

		resp, _ := http.Get(url)
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		for _, b := range body {
			c := strings.ToLower(string(b))
			cIndex := strings.Index(allLetters, c)
			if cIndex >= 0 {
				frequency[cIndex] += 1
				fmt.Println("letter: ", c)
			}

			fmt.Println("Completed: ", url)
			result <- frequency
		}
	}()
	return result
}

// TODO fix because it is not counting letters correctly
func main() {
	results := make([]<-chan []int, 0)
	totalFrequencies := make([]int, 26)

	for i := 1000; i <= 1200; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		results = append(results, countLetters(url))
	}

	for _, c := range results {
		frequencyResult := <-c
		fmt.Println("frequencyResult", frequencyResult)
		for i := 0; i < 26; i++ {
			totalFrequencies[i] += frequencyResult[i]
		}
	}

	for i, c := range allLetters {
		fmt.Printf("%c-%d \n", c, totalFrequencies[i])
	}
}
