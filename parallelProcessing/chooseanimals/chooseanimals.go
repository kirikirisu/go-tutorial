package main

import (
	"fmt"
	"math/rand"
	"time"
)

func choose(cf, cd, quit chan int) {
	for {
		select {
		case count := <-cf:
			fmt.Printf("カエルが%d匹\n", count)
		case count := <-cd:
			fmt.Printf("アヒルが%d匹\n", count)
		case <-quit:
			fmt.Println("終了")
			return
		}
	}
}

func countRandom(cf, cd, quit chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		count := rand.Intn(10) + 1
		if count%2 == 0 {
			cf <- count
		} else {
			cd <- count
		}
	}
	quit <- 1
}

func main() {
	cf := make(chan int)
	cd := make(chan int)
	quit := make(chan int)

	go countRandom(cf, cd, quit)
	choose(cf, cd, quit)
}
