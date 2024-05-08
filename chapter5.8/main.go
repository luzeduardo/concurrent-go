package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	playersInGame := 4
	for playerId := 0; playerId < 4; playerId++ {
		go playerHanlder(cond, &playersInGame, playerId)
		time.Sleep(time.Duration(playerId) * time.Second)
	}

	fmt.Println("Game started!")
}

// Each time a new player joins the game the goroutine checks for the condition to have all players connected before
// waking up the other suspended threads with the Broadcast
func playerHanlder(cond *sync.Cond, playersRemaining *int, playerId int) {
	cond.L.Lock()
	fmt.Println(playerId, ": Connected")

	*playersRemaining--

	if *playersRemaining == 0 {
		cond.Broadcast()
	}

	if *playersRemaining > 0 {
		waitingMessage := fmt.Sprintf("Waiting for more %d players", *playersRemaining)
		fmt.Println(playerId, waitingMessage)
		cond.Wait()
	}

	cond.L.Unlock()
	fmt.Println("All players connected. Ready to play", playerId)
}
