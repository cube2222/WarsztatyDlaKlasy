package main

import (
	"fmt"
	"strconv"
)

func main() {
	myChan := make(chan int)

	go pong(myChan)

	myChan <- 0

	for i := 0; i < 100; i++ {
		item := <- myChan
		fmt.Println(item)
		myChan <- item + 2
	}

	strconv.Atoi()
}

func pong(myChan chan int) {
	for item := range myChan {
		fmt.Println(item)
		myChan <- item + 2
	}
}
