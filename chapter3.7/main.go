package main

import (
	"fmt"
	"time"
)

func stingy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 10
	}
	fmt.Println("Stingy done")
}

func spendy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 10
	}
	fmt.Println("Speny done")
}

func main() {
	money := 100
	go stingy(&money)
	go spendy(&money)
	time.Sleep(2 * time.Second)
	println("Money in banck account: ", money)
}
