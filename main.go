package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"porter/config"
	"porter/db"
	logic "porter/pkg/logic/client"
	"strings"
	"time"

	// gql "porter/pkg/logic/gql"
	"porter/routers"

	"github.com/bitly/go-simplejson"
	"github.com/prometheus/common/log"
)

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}

func setFlag() {
	flag.Usage = usage
	err := flag.Set("alsologtostderr", "true")
	if err != nil {
		log.Error(err)
	}
	err = flag.Set("v", "5")
	if err != nil {
		log.Error(err)
	}
	err = flag.Set("log_dir", "log_folder")
	if err != nil {
		log.Error(err)
	}
	flag.Parse()
}

func init() {
	setFlag()
	//command all if run in saas
	/*
		os.Setenv("MONGODB_URL", "52.187.110.12:27017")
		os.Setenv("MONGODB_DATABASE", "87e1dc58-4c20-4e65-ad81-507270f6bdac")
		os.Setenv("MONGODB_USERNAME", "19e0ce80-af51-404c-8d55-9edefcbd4bdf")
		os.Setenv("MONGODB_PASSWORD", "TYyvTeVemAlJzzuq4w3sBr2D")

		os.Setenv("IFP_DESK_USERNAME", "devanliang@iii.org.tw")
		os.Setenv("IFP_DESK_PASSWORD", "Abcd1234#")
		// os.Setenv("IFP_DESK_API_URL", "https://ifp-organizer-impelex-eks011.hz.wise-paas.com.cn/graphql")

		os.Setenv("IFP_DESK_API_URL", "https://ifp-organizer-testingsa1-eks002.sa.wise-paas.com/graphql") //匯出: 測試環境
		// os.Setenv("IFP_DESK_API_URL", "https://ifp-organizer-testingsa1-eks002.sa.wise-paas.com/graphql") //匯入: 測試環境。
		// os.Setenv("IFP_DESK_API_URL", "https://ifp-organizer-tienkang-eks002.sa.wise-paas.com/graphql") //天岡
		// os.Setenv("IFP_DESK_API_URL", "https://ifp-organizer-training-eks011.hz.wise-paas.com.cn/graphql") //training
		// os.Setenv("IFP_DESK_API_URL", "https://ifp-organizer-impelex-eks011.hz.wise-paas.com.cn/graphql") //匯出: 銳鼎
		// os.Setenv("IFP_DESK_API_URL", "https://ifp-organizer-impelex-eks011.hz.wise-paas.com.cn/graphql") //匯出: 銳鼎

		//for test appSecret
		// os.Setenv("datacenter", "hz")
		// os.Setenv("workspace", "52434e96-f390-474c-8bf1-15e4802fc4fc")
		// os.Setenv("cluster", "eks011")
		// os.Setenv("namespace", "impelex")
		// os.Setenv("SSO_API_URL", "https://api-sso-ensaas.hz.wise-paas.com.cn/v4.0")
	*/

	initGlobalVar()

	if config.IfpDeskUsername != "" && config.IfpDeskPassword != "" {
		logic.Loop_RefreshTokenByUserPwd() //use user pwd
	} else {
		go logic.RefreshTokenByAppSecret() //use app secret
	}

	db.StartMongo()
	go db.MongoHealCheckLoop()
}

func initGlobalVar() {

	config.Datacenter = os.Getenv("datacenter")
	config.Workspace = os.Getenv("workspace")
	config.Cluster = os.Getenv("cluster")
	config.Namespace = os.Getenv("namespace")
	config.AppID = os.Getenv("appID")
	config.ClientName = os.Getenv("clientName")
	config.ServiceName = os.Getenv("serviceName")
	config.AppSecretFile = os.Getenv("IFP_APP_SECRET_FILE")
	external := os.Getenv("external")
	if config.Namespace == "ifpsdev" {
		config.SSOURL = "https://api-sso-ensaas.hz.wise-paas.com.cn/v4.0"
	} else {
		//#fix 改為WISE_PAAS_SSO_API_URL
		config.SSOURL = os.Getenv("WISE_PAAS_SSO_API_URL")
		fmt.Println("[ENV] WISE_PAAS_SSO_API_URL=", config.SSOURL)
		// config.SSOURL = "https://api-sso-ensaas." + external + "/v4.0"
	}

	fmt.Println("[ENV] IFP_DESK_API_URL=", os.Getenv("IFP_DESK_API_URL"))
	ifps_desk_api_url := os.Getenv("IFP_DESK_API_URL")
	if len(ifps_desk_api_url) != 0 {
		//到時候會給過來像 http://ifps-ifp-organizer:80/graphql 或是 https://ifp-organizer-${NAMESPACE}-${CLUSTER}.hz.wise-paas.com.cn/graphql
		ifps_desk_api_url = strings.ReplaceAll(ifps_desk_api_url, "${NAMESPACE}", config.Namespace)
		ifps_desk_api_url = strings.ReplaceAll(ifps_desk_api_url, "${CLUSTER}", config.Cluster)
		ifps_desk_api_url = strings.ReplaceAll(ifps_desk_api_url, "${DATACENTER}", config.Datacenter)
		config.IFPURL = ifps_desk_api_url
		config.IFP_URL = config.IFPURL
		config.IFP_URL_IN = config.IFPURL
	} else {
		config.IFPURL = "https://ifp-organizer-" + config.Namespace + "-" + config.Cluster + "." + external + "/graphql"
		config.IFP_URL = config.IFPURL
		config.IFP_URL_IN = config.IFPURL
	}
	fmt.Println("[ENV] IFP_DESK_API_URL After Replacing=", ifps_desk_api_url)

	config.IfpDeskUsername = os.Getenv("IFP_DESK_USERNAME")
	config.IfpDeskPassword = os.Getenv("IFP_DESK_PASSWORD")

	// ensaas services
	ensaasService := os.Getenv("ENSAAS_SERVICES")
	if len(ensaasService) != 0 {
		// decoded, err := base64.StdEncoding.DecodeString(ensaasService)
		// if err != nil {
		// 	fmt.Println("decode error:", err)
		// 	return
		// }
		// fmt.Println(string(decoded))

		// m, _ := simplejson.NewJson(decoded)
		// mongodb := m.Get("mongodb").GetIndex(0).Get("credentials").MustMap()

		m, _ := simplejson.NewJson([]byte(ensaasService))
		mongodb := m.Get("mongodb").GetIndex(0).Get("credentials").MustMap()
		config.MongodbURL = mongodb["externalHosts"].(string)
		config.MongodbDatabase = mongodb["database"].(string)
		config.MongodbUsername = mongodb["username"].(string)
		config.MongodbPassword = mongodb["password"].(string)
	} else {
		config.MongodbURL = os.Getenv("MONGODB_URL")
		config.MongodbDatabase = os.Getenv("MONGODB_DATABASE")
		config.MongodbUsername = os.Getenv("MONGODB_USERNAME")
		mongodb_password_file := os.Getenv("MONGODB_PASSWORD_FILE")
		config.MongodbPassword = os.Getenv("MONGODB_PASSWORD")
		if mongodb_password_file != "" {
			mongodb_pwd, err := os.ReadFile(mongodb_password_file)
			if err != nil {
				fmt.Println("Read mongodb_password_file error:", err)
			} else {
				config.MongodbPassword = strings.Trim(string(mongodb_pwd), "\n")
			}
		}
		config.MongodbSource = os.Getenv("MONGODB_AUTH_SOURCE")
	}

	fmt.Println("----------", time.Now().In(config.TaipeiTimeZone), "----------")
	fmt.Println("IFP -> URL:", config.IFPURL)
	fmt.Println("SSO -> URL:", config.SSOURL)
	fmt.Println("MongoDB Connect ->", " URL:", config.MongodbURL, " Database:", config.MongodbDatabase)
}

func testFn() {
	// gql.TestQuery()
	// gql.TestMutation()
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
