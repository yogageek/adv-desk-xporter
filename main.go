package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"porter/db"
	logic "porter/pkg/logic/client"
	gql "porter/pkg/logic/gql"
	"porter/routers"
)

// var m *sync.Mutex
// m = new(sync.Mutex)
// m.Lock()
// m.Unlock()

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}

func setFlag() {
	flag.Usage = usage

	//after setting here, no more need to put in .vscode args
	flag.Set("alsologtostderr", "true")
	flag.Set("v", "5")        //Info,Error...這種的不用set v就印的出來, glog.V(3)...這種的要set v才印的出來
	flag.Set("log_dir", "gg") //put glog log data into "gg" folder, (注意要先把資料夾創出來!)
	//stderrthreshold确保了只有大于或者等于该级别的日志才会被输出到stderr中，也就是标准错误输出中，默认为ERROR。当设置为FATAL时候，不会再有任何error信息的输出。

	flag.Parse() //解析上面的set。 after parse(), so that your flag.set start effected
}

func init() {
	// os.Setenv("IFP_URL", "https://ifp-organizer-tienkang-eks002.sa.wise-paas.com/graphql") //天岡
	// os.Setenv("IFP_URL", "https://ifp-organizer-training-eks011.hz.wise-paas.com.cn/graphql") //training

	// os.Setenv("IFP_URL", "https://ifp-organizer-impelex-eks011.hz.wise-paas.com.cn/graphql") //匯出: 銳鼎
	// os.Setenv("IFP_URL_IN", "https://ifp-organizer-impelex-eks011.hz.wise-paas.com.cn/graphql") //匯出: 銳鼎

	os.Setenv("IFP_URL", "https://ifp-organizer-testingsa1-eks002.sa.wise-paas.com/graphql")    //匯出: 測試環境
	os.Setenv("IFP_URL_IN", "https://ifp-organizer-testingsa1-eks002.sa.wise-paas.com/graphql") //匯入: 測試環境。

	os.Setenv("MONGODB_URL", "52.187.110.12:27017")
	os.Setenv("MONGODB_DATABASE", "87e1dc58-4c20-4e65-ad81-507270f6bdac")
	os.Setenv("MONGODB_USERNAME", "19e0ce80-af51-404c-8d55-9edefcbd4bdf")
	os.Setenv("MONGODB_PASSWORD", "TYyvTeVemAlJzzuq4w3sBr2D")

	setFlag()
	db.StartMongo()
	go db.MongoHealCheckLoop()
	logic.LoopRefreshToken()

	// runtime.GOMAXPROCS(1)

}

func testFn() {
	// gql.TestQuery()
	gql.TestMutation()
	//測試export
	// controller.Export()
}
func main() {
	testFn()
	startServer()
}

func startServer() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: router,
		// ReadTimeout: time.Second * 120,
		// WriteTimeout:      time.Second * 120,
		// MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

// func mywebsocket() {
// 	upgrader := &websocket.Upgrader{
// 		//如果有 cross domain 的需求，可加入這個，不檢查 cross domain
// 		CheckOrigin: func(r *http.Request) bool { return true },
// 	}
