package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateAmounts(n int) <-chan int {
	amounts := make(chan int)
	go func() {
		defer close(amounts)
		for i := 0; i < n; i++ {
			amounts <- rand.Intn(100) + 1
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return amounts
}

// tis pattern of merging channel data in one stream is referred as fan-in pattern.
// using a select to merge a different source only works when there is a fixed number of sources
func main() {
	sales := generateAmounts(50)
	expenses := generateAmounts(40)
	endOfDayAmount := 0
	// reads until both of channels are nil to proceed with another flow of the main goroutine
	for sales != nil || expenses != nil {
		// select keeps reading from a closed channel but it will return a default type value. in this case 0
		// so to avoid this read the return flag to check if it still has data otherwise assign a nil
		// select will disable that specifi select statement
		select {
		case sale, moreData := <-sales:
			if moreData {
				fmt.Println("Sale of: ", sale)
				endOfDayAmount += sale
			} else {
				sales = nil
			}
		case expense, moreData := <-expenses:
			if moreData {
				fmt.Println("Expense of: ", expense)
				endOfDayAmount -= expense

			} else {
				expenses = nil
			}
		}
	}

	fmt.Println("End of the day profit and loss: ", endOfDayAmount)
}
