package main

import (
	"fmt"
	"net/http"
	"os"
	logic "porter/pkg/logic/client"
	"porter/routers"
	"time"
)

// var m *sync.Mutex
// m = new(sync.Mutex)
// m.Lock()
// m.Unlock()

func init() {
	// var IFP_URL = "https://ifp-organizer-training-eks011.hz.wise-paas.com.cn/graphql"
	// var IFP_URL = "https://ifp-organizer-tienkang-eks002.sa.wise-paas.com/graphql" //天岡
	os.Setenv("IFP_URL", "https://ifp-organizer-impelex-eks011.hz.wise-paas.com.cn/graphql")    //匯出: 銳鼎
	os.Setenv("IFP_URL_IN", "https://ifp-organizer-testingsa1-eks002.sa.wise-paas.com/graphql") //匯入: 測試環境。
	logic.LoopRefreshToken()

	// runtime.GOMAXPROCS(1)
}

func main() {
	// logic.Export()
	// logic.Import()
	// logic.QueryMachines()
	// logic.AddMachineSample()
	// logic.QueryGroups()
	// qp := logic.QueryParameters("TWFjaGluZQ.X0OAStZ5cgAG6yN-", "") //cursor:RGF0ZQ.MTU5OTY0MDA2MDAxOQ.X1iR-4RZEgAGpSQj
	// util.PrintBlue(qp)
	// logic.AddParameterSample()

	startServer()
}

func startServer() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:        fmt.Sprintf(":%d", 8080),
		Handler:     router,
		ReadTimeout: time.Second * 120,
		// WriteTimeout:      time.Second * 120,
		// MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

//接下來掛server 3支api

//先測匯出是否可以導出一個檔案

//在測匯入是否可以選擇一個檔案

//最後做取完成狀態率( long polling API，可以讓所有使用者能夠即時知道現在是否有匯入匯出的工作正在做。)

// func mywebsocket() {
// 	upgrader := &websocket.Upgrader{
// 		//如果有 cross domain 的需求，可加入這個，不檢查 cross domain
// 		CheckOrigin: func(r *http.Request) bool { return true },
// 	}

// 	http.HandleFunc("/websocket/config/file/status", func(w http.ResponseWriter, r *http.Request) {
// 		//web"開啟連接"後就會一路進來

// 		//step1: set upgrader
// 		c, err := upgrader.Upgrade(w, r, nil)
// 		if err != nil {
// 			log.Println("upgrade:", err)
// 			return
// 		}
// 		defer func() {
// 			log.Println("disconnect !!")
// 			c.Close()
// 		}()

// 		//setp2: do your logic
// 		//測試(如果正在做)直接先write
// 		res := logic.Res
// 		if res.State == logic.StateDoing {
// 			rtn := func() []byte {
// 				myR := map[string]interface{}{
// 					"error": "still in progress",
// 					"state": int(logic.StateDoing),
// 				}
// 				b, _ := json.MarshalIndent(myR, "", " ")
// 				return b
// 			}()
// 			err = c.WriteMessage(1, rtn)
// 			if err != nil {
// 				log.Println("write:", err)
// 			}
// 		}

// 		for { //這個for是常規寫法一定要加

// 			//如果有read,才做write(前端來問進度才返回給他)
// 			mtype, msg, err := c.ReadMessage() //web"開啟連接"後就會一路進來並停在這, 如果web有發消息就會往下走並再回來
// 			if err != nil {
// 				log.Println("read:", err) //web關閉連接後跳出
// 				break
// 			}
// 			log.Printf("receive: %s\n", msg)

// 			//只需將當前status狀態寫到這裡(chanel寫法)
// 			//---
// 			res := logic.Res
// 			msg, _ = json.MarshalIndent(res, "", " ")
// 			//---
// 			err = c.WriteMessage(mtype, msg)
// 			if err != nil {
// 				log.Println("write:", err)
// 				break
// 			}
// 		}

// 	})
// 	// log.Println("server start at :8899")
// 	// log.Fatal(http.ListenAndServe(":8899", nil))

// 	log.Println("server start at :8000")
// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }
