package method

import (
	"fmt"
)

//加在外面跟裡面有何差異?
var C = make(chan int) // Allocate a channel.

func ChannelCount(s string, i int) {
	go func() { //避免 c<-1 卡住時導致外層卡死
		C <- i
		fmt.Println(s, ":", i)
	}()
}

// style1
func ChannelGetCount1() {
	go func() {
		for {
			i := <-C
			fmt.Println("take out i:", i)
		}
	}()
}

// style2
func ChannelGetCount2() (i int) {
	i = <-C //會卡在這
	fmt.Println("take out i:", i)
	return
}
