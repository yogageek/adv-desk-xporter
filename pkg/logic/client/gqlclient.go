package logic

import (
	"net/http"
	"net/http/cookiejar"
	"porter/config"

	"github.com/shurcooL/graphql"
)

var (
	GclientQ *graphql.Client
	GclientM *graphql.Client
)

func PrepareGQLCLient() {
	if config.AdminUsername != "" && config.AdminPassword != "" {
		PrepareGQLClientByUserPwd()
	} else {
		PrepareGQLClientByAppSecret()
	}
}

func PrepareGQLClientByAppSecret() {
	NewGQLClientHeader1()
	NewGQLClientHeader2()
}

func NewGQLClientHeader1() {
	httpClient := http.DefaultClient
	rt := WithHeader(httpClient.Transport)
	rt.Set("X-Ifp-App-Secret", config.Token)
	httpClient.Transport = rt
	GclientQ = graphql.NewClient(config.IFP_URL, httpClient)

	//-------test-------
	// GclientQ = graphql.NewClient("https://ifp-organizer-impelex-eks011.hz.wise-paas.com.cn/graphql", httpClient)

	// gqlQuery := model.QueryMachineStatuses

	// variables := map[string]interface{}{
	// 	"layer1Only": graphql.Boolean(true),
	// }

	// err := GclientQ.Query(context.Background(), &gqlQuery, variables)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(gqlQuery)
}

func NewGQLClientHeader2() {
	httpClient := http.DefaultClient
	rt := WithHeader(httpClient.Transport)
	rt.Set("X-Ifp-App-Secret", config.Token)
	httpClient.Transport = rt
	GclientM = graphql.NewClient(config.IFP_URL_IN, httpClient)
}

func PrepareGQLClientByUserPwd() {
	NewGQLClientCookie1()
	NewGQLClientCookie2()
}

func NewGQLClientCookie1() {
	var cookies []*http.Cookie
	var cookieJar *cookiejar.Jar
	cookieJar, _ = cookiejar.New(nil)
	cookies = nil
	httpClient := &http.Client{
		Jar: cookieJar, //must put
	}

	//------------->
	//handling cookie
	req, _ := http.NewRequest("GET", config.IFP_URL, nil)
	req.Header.Set("cookie", UserPwdToken) // set cookie by req (better way)
	// cookies = cookieJar.Cookies(req.URL) // not good way
	cookies = req.Cookies()
	httpClient.Jar.SetCookies(req.URL, cookies)

	//set graphql client for query
	GclientQ = graphql.NewClient(config.IFP_URL, httpClient)
}

func NewGQLClientCookie2() {
	var cookies []*http.Cookie
	var cookieJar *cookiejar.Jar
	cookieJar, _ = cookiejar.New(nil)
	cookies = nil
	httpClient := &http.Client{
		Jar: cookieJar, //must put
	}

	//------------->
	//handling cookie
	req, _ := http.NewRequest("GET", config.IFP_URL_IN, nil)
	req.Header.Set("cookie", UserPwdToken2) // set cookie by req (better way)
	// cookies = cookieJar.Cookies(req.URL) // not good way
	cookies = req.Cookies()
	httpClient.Jar.SetCookies(req.URL, cookies)

	//set graphql client for mutation
	GclientM = graphql.NewClient(config.IFP_URL_IN, httpClient)
}
