package logic

import (
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
var TheStructChan LogChan

type LogChan struct {
	ChanS    chan ChanStruct
	Err      interface{}
	ErrCount int
}

type ChanStruct struct {
	Name string `json:"errName,omitempty"`
	Msg  string `json:"errMsg,omitempty"`
}

func (c *LogChan) InitChan() {
	c.ChanS = make(chan ChanStruct)
}

//跨PKG設置*chan
func SetChan(c *LogChan) {
	TheStructChan = *c
}

//跨PKG傳送*chan
func GetChan() *LogChan {
	return &TheStructChan
}

func (c *LogChan) PutChan(err error, name string) {
	c.ChanS <- ChanStruct{
		Name: name,
		Msg:  err.Error(),
	}
}

func (c *LogChan) TakeChan() {
	var cslist []ChanStruct
	for {
		cs, ok := <-c.ChanS
		if ok {
			cslist = append(cslist, cs)
			c.ErrCount = len(cslist)
			c.Err = cslist
		}
	}
}

func sample() {
	sc1 := &LogChan{
		// Name: "created with Channel",
		// oldCh: make(chan error, 1),
	}

	sc2 := &LogChan{
		// Name: "created without channel",
	}
	sc2.InitChan()

	var err error
	go func() {
		sc1.PutChan(err, "A")
	}()
	go func() {
		sc2.PutChan(err, "B")
	}()
	time.Sleep(5 * time.Second)

}
