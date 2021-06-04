package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"porter/config"
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

func initGlobalVar() {

	/*
		### 地端
		- 讀取 `IFP_APP_SECRET_FILE` 路徑檔案的內容

		### 雲端
		- 環境變數會有
		- WISE_PAAS_SERVICE_NAME
		- WISE_PAAS_SSO_API_URL
		- namespace
		- appID
		- datacenter
		- workspace
		- cluster
	*/
	config.Datacenter = os.Getenv("datacenter")
	config.Workspace = os.Getenv("workspace")
	config.Cluster = os.Getenv("cluster")
	config.Namespace = os.Getenv("namespace")
	config.AppID = os.Getenv("appID")
	if config.Namespace == "ifpsdev" || config.Namespace == "ifpsdemo" {
		config.SSOURL = "https://api-sso-ensaas.hz.wise-paas.com.cn/v4.0"
	} else {
		config.SSOURL = os.Getenv("SSO_API_URL")
	}
	external := os.Getenv("external")

	ifps_desk_api_url := os.Getenv("IFP_DESK_API_URL")
	if len(ifps_desk_api_url) != 0 {
		config.IFPURL = ifps_desk_api_url
	} else {
		config.IFPURL = "https://ifp-organizer-" + config.Namespace + "-" + config.Cluster + "." + external + "/graphql"
	}

	config.AdminUsername = os.Getenv("IFP_DESK_USERNAME")
	config.AdminPassword = os.Getenv("IFP_DESK_PASSWORD")
}

func init() {
	// var IFP_URL = "https://ifp-organizer-tienkang-eks002.sa.wise-paas.com/graphql" //天岡
	// os.Setenv("IFP_URL", "https://ifp-organizer-training-eks011.hz.wise-paas.com.cn/graphql")
	os.Setenv("IFP_URL", "https://ifp-organizer-impelex-eks011.hz.wise-paas.com.cn/graphql") //匯出: 銳鼎
	// os.Setenv("IFP_URL_IN", "https://ifp-organizer-testingsa1-eks002.sa.wise-paas.com/graphql") //匯入: 測試環境。
	os.Setenv("IFP_URL_IN", "https://ifp-organizer-impelex-eks011.hz.wise-paas.com.cn/graphql") //匯出: 銳鼎
	os.Setenv("MONGODB_URL", "52.187.110.12:27017")
	os.Setenv("MONGODB_DATABASE", "87e1dc58-4c20-4e65-ad81-507270f6bdac")
	os.Setenv("MONGODB_USERNAME", "19e0ce80-af51-404c-8d55-9edefcbd4bdf")
	os.Setenv("MONGODB_PASSWORD", "TYyvTeVemAlJzzuq4w3sBr2D")

	// 2021/05/31
	os.Setenv("SSO_API_URL", "https://api-sso-ensaas.sa.wise-paas.com/v4.0")
	os.Setenv("datacenter", "sa")
	os.Setenv("workspace", "4d468e6a-0744-40f0-8c05-e89c31440fc3")
	os.Setenv("cluster", "eks005")
	os.Setenv("namespace", "ifpsdemo")
	os.Setenv("appID", "IunAtiS_vsxaFcgN5QluNwfxbXN7tUu3-1614837178")
	os.Setenv("IFP_DESK_USERNAME", "devanliang@iii.org.tw")
	os.Setenv("IFP_DESK_PASSWORD", "Abcd1234#")
	os.Setenv("IFP_DESK_API_URL", "https://ifp-organizer-impelex-eks011.hz.wise-paas.com.cn/graphql")
	initGlobalVar()
	go logic.RefreshTokenByAppSecret()
	// MDVlOGM1N2QtYzQ0MC0xMWViLTlmNzUtYmUxNGNlMTk2ZDQ2
	// 2021/05/31 End

	setFlag()
	db.StartMongo()
	go db.MongoHealCheckLoop()
	logic.LoopRefreshToken()

	// runtime.GOMAXPROCS(1)

}

func testFn() {
	gql.TestQuery()
	//測試export
	// controller.Export()
}
func main() {
	gql.Run()
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
