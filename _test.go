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
	for i := 0; i < 10; i++ {
		c := make(chan int)
		go worker(i, c)
		c <- 1
		c <- 2
	}
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
