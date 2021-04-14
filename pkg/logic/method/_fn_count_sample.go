// package method
package main

import (
	"fmt"
	"time"
)

//加在外面跟裡面有何差異?
var c = make(chan int) // Allocate a channel.

func Count(s string) {

	// c := make(chan int)
	// go func() {
	// 	close(c)
	// }()
	// <-c

	// c := make(chan int) // Allocate a channel.
	// go func() {
	c <- 1
	// }()
	fmt.Println(s, "count done")
}

func GetChannelCount() (i int) {
	/*
		go func() {
			for {
				n := <-c
				fmt.Println(n)
			}
		}()
	*/

	go func() {
		for {
			result := <-c
			i = i + result
			fmt.Println("i:", i)
			if i == 100 {
				break
			}
		}
	}()

	return i
}

func main() {
	go func() {
		for i := 0; i < 100; i++ {
			Count("a")
		}
	}()

	GetChannelCount()

	time.Sleep(2 * time.Second)

	go func() {
		for i := 101; i < 110; i++ {
			Count("b")
		}
	}()

	fmt.Println("yes")

	time.Sleep(10 * time.Second)
}

//以上範例可觀察 送了沒取(一直送) 或取了沒送(一直取) 是如何造成deadlock
//如果是放在go func裡不會噴error 但會卡死
//如果沒放在go func會直接噴error

//因為channel是gorouting到gorouting的交互，我們必須要採用另外一個gorouting去收他
//這發了一個數據沒人收是會deadlock(死鎖)的,因為我們需要有一個人來負責收
