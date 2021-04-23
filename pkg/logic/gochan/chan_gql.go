package logic

import (
	"fmt"
	"time"
)

/*
PKG gochan init global chan
PKG logic(import_controller) take out chan(方法來自於PKG gochan)
PKG logic(import_api...) send data to chan(方法來自於PKG gochan)


PKG router init chan，把 *chan 放入方法的參數傳送到其他 PKG
PKG gql，new func(接收*chan) 並塞值進去
*/

//跨PKG使用共同*chan
var TheStructChan StructChan

type StructChan struct {
	Name  string
	Sum   int
	Errs  []string
	errCh chan error
}

func (c *StructChan) InitChan() {
	c.errCh = make(chan error, 10000)
}

//跨PKG設置*chan
func SetChan(sc *StructChan) {
	TheStructChan = *sc
}

//跨PKG傳送*chan
func GetChan() *StructChan {
	return &TheStructChan
}

func (sc *StructChan) SendToChan(err error) {
	sc.errCh <- err
	fmt.Println(sc.Name, "SendToChan ok")
}

func (sc *StructChan) TakeFromChan() {
	a := 0
	for {
		a++
		err := <-sc.errCh
		fmt.Println("TakeFromChan err:", err.Error())
		sc.Errs = append(sc.Errs, err.Error())
		if a > 5 {
			break
		}
	}
}

func main() {
	sc1 := &StructChan{
		Name:  "created with Channel",
		errCh: make(chan error, 1),
	}

	sc2 := &StructChan{
		Name: "created without channel",
	}
	sc2.InitChan()

	var err error
	go func() {
		sc1.SendToChan(err)
	}()
	go func() {
		sc2.SendToChan(err)
	}()
	time.Sleep(5 * time.Second)

}
