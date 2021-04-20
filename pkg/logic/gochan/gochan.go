package logic

import (
	. "porter/pkg/logic/vars"
	"time"
)

/*

boool := gochan.ChannelOut() //如果匯入資料送完 這裡取完 會停在這行  only mStatus

		// if ChannelOut()==true 代表有從channel取出值
		if boool {
			snedEventProcess()
		} else {
			sendEventDone()
			vars.Res.State = vars.StateDone
			// break //跳出迴圈 重新監測事件
			PrintJson(vars.Res.Details)

			//debugging 暫時 不然多線程關不掉
			// return //return 等於斷開連結(不可!)

			goto case2

		}
		// testing write ws
		// s := strconv.Itoa(i)
		// ws.WriteMessage(1, []byte(s))
		goto case2

*/

// var Cm = make(chan map[string]int)
var channelA = make(chan int)
var channelB = make(chan int)
var channelC = make(chan int)
var channelD = make(chan int)
var channelE = make(chan int)
var channelF = make(chan int)

//放入
func ChannelIn(s string, i int) {
	switch s {
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
			SetResponseCount(MachineStatus, i)
			return true
		case i := <-channelB:
			// fmt.Println(i)
			SetResponseCount(MappingRule, i)
			return true
		case i := <-channelC:
			// fmt.Println(i)
			SetResponseCount(Profile, i)
			return true
		case i := <-channelD:
			// fmt.Println(i)
			SetResponseCount(Group, i)
			return true
		case i := <-channelE:
			// fmt.Println(i)
			SetResponseCount(Machine, i)
			return true
		case i := <-channelF:
			// fmt.Println(i)
			SetResponseCount(Parameter, i)
			return true
		case <-time.After(5 * time.Second):
			break
			// default:
			// 	break
		}
		return false
	}

	//也可用for range的方式取出channel
	// var ints []int
	// for i := range channelA {
	// 	ints = append(ints, i)
	// 	fmt.Println("A:", ints)
	// 	break
	// }

}

// func ChannelGetCount1() {
// 	go func() {
// 		for {
// 			i := <-C
// 			fmt.Println("take out i:", i)
// 		}
// 	}()
// }
