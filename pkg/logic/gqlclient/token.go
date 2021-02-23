package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
)

var (
	Token string

	ADMIN_USERNAME    = "devanliang@iii.org.tw"
	ADMIN_PASSWORD    = "Abcd1234#"
	IFP_URL           = "https://ifp-organizer-training-eks011.hz.wise-paas.com.cn/graphql"
	TaipeiTimeZone, _ = time.LoadLocation("Asia/Taipei")

	MONGODB_URL      = "52.187.110.12:27017"
	MONGODB_DATABASE = "ifp-data-hub-dev"
	MONGODB_USERNAME = "e270673c-ce08-4c35-93e2-333ed103736f"
	MONGODB_PASSWORD = "VUSkt9bbTKSTzb7ZArp36jLk"
)

func RefreshToken() {
	// for {
	fmt.Println("----------", time.Now().In(TaipeiTimeZone), "----------")
	fmt.Println("refreshToken")
	httpClient := &http.Client{}
	content := map[string]string{"username": ADMIN_USERNAME, "password": ADMIN_PASSWORD}
	variable := map[string]interface{}{"input": content}
	httpRequestBody, _ := json.Marshal(map[string]interface{}{
		"query":     "mutation signIn($input: SignInInput!) {   signIn(input: $input) {     user {       name       __typename     }     __typename   } }",
		"variables": variable,
	})
	request, _ := http.NewRequest("POST", IFP_URL, bytes.NewBuffer(httpRequestBody))
	request.Header.Set("Content-Type", "application/json")
	response, _ := httpClient.Do(request)
	m, _ := simplejson.NewFromReader(response.Body)
	for {
		if len(m.Get("errors").MustArray()) == 0 {
			break
		} else {
			httpRequestBody, _ = json.Marshal(map[string]interface{}{
				"query":     "mutation signIn($input: SignInInput!) {   signIn(input: $input) {     user {       name       __typename     }     __typename   } }",
				"variables": variable,
			})
			request, _ = http.NewRequest("POST", IFP_URL, bytes.NewBuffer(httpRequestBody))
			request.Header.Set("Content-Type", "application/json")
			response, _ = httpClient.Do(request)
			m, _ = simplejson.NewFromReader(response.Body)
		}
	}
	// fmt.Println("-- GraphQL API End", time.Now().In(taipeiTimeZone))
	header := response.Header
	// fmt.Println(header)
	// m, _ := simplejson.NewFromReader(response.Header)
	cookie := header["Set-Cookie"]
	tempSplit := strings.Split(cookie[0], ";")
	ifpToken := tempSplit[0]
	tempSplit = strings.Split(cookie[1], ";")
	eiToken := tempSplit[0]
	Token = ifpToken + ";" + eiToken
	// fmt.Println("Token:", Token)
	os.Setenv("Token", Token)
	// time.Sleep(60 * time.Minute)
	// }
}
