package main

import "fmt"

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

func main() {
	quit := make(chan int)
	defer close(quit)
	fmt.Println("a")
	results := generateUrls(quit)
	fmt.Println("b")

	for result := range results {
		fmt.Println("f")
		fmt.Println(result)
	}

	fmt.Println("g")
}
