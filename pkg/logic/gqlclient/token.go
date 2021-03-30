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
	ADMIN_USERNAME = "devanliang@iii.org.tw"
	ADMIN_PASSWORD = "Abcd1234#"

	MONGODB_URL      = "52.187.110.12:27017"
	MONGODB_DATABASE = "ifp-data-hub-dev"
	MONGODB_USERNAME = "e270673c-ce08-4c35-93e2-333ed103736f"
	MONGODB_PASSWORD = "VUSkt9bbTKSTzb7ZArp36jLk"

	TaipeiTimeZone, _ = time.LoadLocation("Asia/Taipei")
)

var (
	Token  string
	Token2 string
)

var (
	IFP_URL    string
	IFP_URL_IN string
)

func DoRefreshToken() {
	IFP_URL = os.Getenv("IFP_URL")
	Token = RefreshToken(IFP_URL)
	// fmt.Println(Token)

	IFP_URL_IN = os.Getenv("IFP_URL_IN")
	Token2 = RefreshToken(IFP_URL_IN)
	// fmt.Println(Token2)
}

func RefreshToken(url string) (token string) {
	// for {
	fmt.Println("----------", time.Now().In(TaipeiTimeZone), "----------")
	fmt.Println("RefreshToken...")
	httpClient := &http.Client{}
	content := map[string]string{"username": ADMIN_USERNAME, "password": ADMIN_PASSWORD}
	variable := map[string]interface{}{"input": content}
	httpRequestBody, _ := json.Marshal(map[string]interface{}{
		"query":     "mutation signIn($input: SignInInput!) {   signIn(input: $input) {     user {       name       __typename     }     __typename   } }",
		"variables": variable,
	})
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(httpRequestBody))
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
			request, _ = http.NewRequest("POST", url, bytes.NewBuffer(httpRequestBody))
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
	token = ifpToken + ";" + eiToken
	return token
	// fmt.Println("Token:", Token)
	// os.Setenv("Token", Token)
	// time.Sleep(60 * time.Minute)
	// }
}
