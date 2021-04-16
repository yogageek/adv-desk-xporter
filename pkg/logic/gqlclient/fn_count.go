package logic

import (
	"fmt"
	. "porter/pkg/logic/var"
	"sync"
)

// var Cm = make(chan map[string]int)

var a = make(chan int)
var b = make(chan int)
var c = make(chan int)
var d = make(chan int)
var e = make(chan int)
var f = make(chan int)

var A int
var B int
var C int
var D int
var E int
var F int

var mux sync.Mutex

//放入
func ChannelCount(s string, i int) {

	var wg sync.WaitGroup
	var mux sync.Mutex

	wg.Add(1)

	// fmt.Println(s, ":", i)
	go func() { //避免 x<-1 卡住時導致外層卡死
		switch s {
		case MachineStatus:
			mux.Lock()
			fmt.Println("正常印出", i)
			a <- i                       //這行有問題
			fmt.Println("~~~~~~~~~~~~~") //要把ws端打開才會往下執行 因為當a<-i沒有取出的話馬上就會卡住
			mux.Unlock()

		case MappingRule:
			fmt.Println("+++++++++++++", i)
			b <- i
		case Profile:
			c <- i
		case Group:
			d <- i
		case Machine:
			e <- i
		case Parameter:
			f <- i
		}
		defer wg.Done()
	}()
	wg.Wait()
}

// style1
// func ChannelGetCount1() {
// 	go func() {
// 		for {
// 			i := <-C
// 			fmt.Println("take out i:", i)
// 		}
// 	}()
// }

var ii []int

//取出
// style2
func ChannelGetCount2() {

	var i int
	var s string
	fmt.Println("wait get channel................")
	select {
	case <-a:
		i = <-a
		ii = append(ii, i) //會短缺
		fmt.Println(ii)
		// if i == 19 {
		// 	fmt.Println(i)
		// }
		s = MachineStatus
		goto setRes

	case <-b:
		i = <-b

		fmt.Println(i)

		s = MappingRule
		goto setRes
	case <-c:
		i = <-c
		s = Profile
		goto setRes
	case <-d:
		i = <-d
		s = Group

	case <-e:
		i = <-e
		s = Machine

	case <-f:
		i = <-f
		s = Parameter

		// default:
		// 	fmt.Println("exit--------------")
		// 這裡不能加defualt 否則因為當沒channel可取後 會結束這個func 讓websocket無限迴圈一直傳
	}

setRes:
	// fmt.Println("set res:", s, i)
	SetResponse(s, i)
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
