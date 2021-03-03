package logic

import (
	"net/http"
	"net/http/cookiejar"

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

var (
	gclientQ *graphql.Client
	gclientM *graphql.Client
)

func NewGQLClient() {
	var cookies []*http.Cookie
	var cookieJar *cookiejar.Jar
	cookieJar, _ = cookiejar.New(nil)
	cookies = nil
	httpClient := &http.Client{
		Jar: cookieJar, //must put
	}

	//------------->
	//handling cookie
	req, _ := http.NewRequest("GET", IFP_URL, nil)
	req.Header.Set("cookie", Token) // set cookie by req (better way)
	// cookies = cookieJar.Cookies(req.URL) // not good way
	cookies = req.Cookies()
	httpClient.Jar.SetCookies(req.URL, cookies)

	//set graphql client for query
	gclientQ = graphql.NewClient(IFP_URL, httpClient)
}

func NewGQLClient2() {
	var cookies []*http.Cookie
	var cookieJar *cookiejar.Jar
	cookieJar, _ = cookiejar.New(nil)
	cookies = nil
	httpClient := &http.Client{
		Jar: cookieJar, //must put
	}

	//------------->
	//handling cookie
	req, _ := http.NewRequest("GET", IFP_URL_IN, nil)
	req.Header.Set("cookie", Token2) // set cookie by req (better way)
	// cookies = cookieJar.Cookies(req.URL) // not good way
	cookies = req.Cookies()
	httpClient.Jar.SetCookies(req.URL, cookies)

	//set graphql client for mutation
	gclientM = graphql.NewClient(IFP_URL_IN, httpClient)
}
