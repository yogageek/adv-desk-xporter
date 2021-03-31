package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		//这里Println改成Printf
		fmt.Printf("Worker %d received %d\n", id, <-c)
	}
}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
