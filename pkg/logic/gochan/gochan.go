package logic

import (
	. "porter/pkg/logic/vars"
	"time"
)

// var Cm = make(chan map[string]int)
var channelA = make(chan int)
var channelB = make(chan int)
var channelC = make(chan int)
var channelD = make(chan int)
var channelE = make(chan int)
var channelF = make(chan int)

//放入
func ChannelIn(name string, i int) {
	switch name {
	case MachineStatus:
		// fmt.Println("A正常印出", i)
		go func(channel chan<- int, order int) {
			channel <- order
		}(channelA, i)
	case MappingRule:
		// fmt.Println("B正常印出", i)
		go func(channel chan<- int, order int) {
			channel <- order
		}(channelB, i)
	case Profile:
		// fmt.Println("C正常印出", i)
		go func(channel chan<- int, order int) {
			channel <- order
		}(channelC, i)
	case Group:
		// fmt.Println("D正常印出", i)
		go func(channel chan<- int, order int) {
			channel <- order
		}(channelD, i)
	case Machine:
		// fmt.Println("E正常印出", i)
		go func(channel chan<- int, order int) {
			channel <- order
		}(channelE, i)
	case Parameter:
		// fmt.Println("F正常印出", i)
		go func(channel chan<- int, order int) {
			channel <- order
		}(channelF, i)
	}
}

//取出
func ChannelOut() bool {
	for {
		select {
		case i := <-channelA:
			// fmt.Println(i)
			Update_PuclicRes_Detail(MachineStatus, i)
			return true
		case i := <-channelB:
			// fmt.Println(i)
			Update_PuclicRes_Detail(MappingRule, i)
			return true
		case i := <-channelC:
			// fmt.Println(i)
			Update_PuclicRes_Detail(Profile, i)
			return true
		case i := <-channelD:
			// fmt.Println(i)
			Update_PuclicRes_Detail(Group, i)
			return true
		case i := <-channelE:
			// fmt.Println(i)
			Update_PuclicRes_Detail(Machine, i)
			return true
		case i := <-channelF:
			// fmt.Println(i)
			Update_PuclicRes_Detail(Parameter, i)
			return true
		case <-time.After(5 * time.Second):
			return false
			// default:
			// 	break
		}
	}
}

//也可用for range的方式取出channel
// var ints []int
// for i := range channelA {
// 	ints = append(ints, i)
// 	fmt.Println("A:", ints)
// 	break
// }

// func ChannelGetCount1() {
// 	go func() {
// 		for {
// 			i := <-C
// 			fmt.Println("take out i:", i)
// 		}
// 	}()
// }
