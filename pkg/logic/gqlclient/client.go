package logic

import (
	"net/http"
	"net/http/cookiejar"
	"os"

	"github.com/shurcooL/graphql"
)

/*
	//set token by oauth2
	src := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: os.Getenv("Token")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
*/
//setting cookies for httpclient
// cookie := &http.Cookie{
// 	Name:  "token",
// 	Value: token,
// 	// Path:  "/",
// 	// Domain: ".weibo.cn",
// }
// cookies = append(cookies, cookie)
/*
	//fail , will get null
	func() {
		b, err := json.Marshal(&Query)
		if err != nil {
			glog.Error(err)
		}
		fmt.Println(string(b))
	}()
*/

// var IFP_URL = "https://ifp-organizer-training-eks011.hz.wise-paas.com.cn/graphql"
// var IFP_URL = "https://ifp-organizer-tienkang-eks002.sa.wise-paas.com/graphql" //天岡
// var IFP_URL = "https://ifp-organizer-impelex-eks011.hz.wise-paas.com.cn/graphql" //匯出: 銳鼎
var IFP_URL = "https://ifp-organizer-testingsa1-eks002.sa.wise-paas.com/graphql" //匯入: 測試環境

var gclient *graphql.Client

func InitGqlClientAndToken() {
	Token := os.Getenv("Token")

	var cookies []*http.Cookie
	var cookieJar *cookiejar.Jar
	cookieJar, _ = cookiejar.New(nil)
	cookies = nil

	httpClient := &http.Client{
		Jar: cookieJar, //must put
	}

	//handling cookie
	req, _ := http.NewRequest("GET", IFP_URL, nil)
	req.Header.Set("cookie", Token) // set cookie by req (better way)
	// cookies = cookieJar.Cookies(req.URL) // not good way
	cookies = req.Cookies()
	httpClient.Jar.SetCookies(req.URL, cookies)

	//set graphql client
	gclient = graphql.NewClient(IFP_URL, httpClient)
}
