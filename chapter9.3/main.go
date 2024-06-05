package main

import (
	"fmt"
	"io"
	"net/http"
)

func generateUrls(quit <-chan int) <-chan string {
	urls := make(chan string)
	go func() {
		defer close(urls)
		for i := 100; i <= 105; i++ {
			url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
			select {
			case urls <- url:
				fmt.Println("c")
			case <-quit:
				fmt.Println("d--")
				return
			}
		}
	}()
	fmt.Println("e")
	// output channel
	return urls
}

func downloadingPages(quit <-chan int, urls <-chan string) <-chan string {
	// receives an input channel and returns an output channel
	pages := make(chan string)
	go func() {
		defer close(pages)

		moreData, url := true, ""
		for moreData {
			select {
			case url, moreData = <-urls:
				resp, _ := http.Get(url)
				if resp.StatusCode != 200 {
					panic("Server error: " + resp.Status)
				}
				body, _ := io.ReadAll(resp.Body)
				pages <- string(body)
				resp.Body.Close()
			case <-quit:
				return
			}
		}
	}()

	return pages
}

func main() {
	quit := make(chan int)
	defer close(quit)

	results := downloadingPages(quit, generateUrls(quit))
	fmt.Println("b")

	for result := range results {
		fmt.Println("f")
		fmt.Println(result)
	}

	fmt.Println("g")
}
