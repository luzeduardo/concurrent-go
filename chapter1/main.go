package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go doWork(i)
	}
	time.Sleep(time.Second * 2)
}

func doWork(id int) {
	fmt.Println(id, "Work started at: ", time.Now().Format("15:04:00"))
	time.Sleep(time.Second * 1)
	fmt.Println(id, "Work finished at: ", time.Now().Format("15:04:00"))
}
