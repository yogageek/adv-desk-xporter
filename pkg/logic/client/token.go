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
	"github.com/prometheus/common/log"
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

func LoopRefreshToken() {

	IFP_URL = os.Getenv("IFP_URL")
	IFP_URL_IN = os.Getenv("IFP_URL_IN")

	go func() {
		for {
			fmt.Println("<<< Loop refresh token >>>")
			var err error
			Token, err = RefreshToken(IFP_URL)
			if err != nil {
				log.Error("RefreshToken IFP_URL fail")
				panic(err)
			}
			Token2, err = RefreshToken(IFP_URL_IN)
			if err != nil {
				log.Error("RefreshToken IFP_URL_IN fail")
				panic(err)
			}
			time.Sleep(15 * time.Minute)
		}
	}()

	// fmt.Println(Token)
	// fmt.Println(Token2)
}

func RefreshToken(url string) (token string, err error) {

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
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(httpRequestBody))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := httpClient.Do(request)
	if err != nil {
		return "", err
	}
	m, _ := simplejson.NewFromReader(response.Body)
	for {
		if len(m.Get("errors").MustArray()) == 0 {
			break
		} else {
			httpRequestBody, _ = json.Marshal(map[string]interface{}{
				"query":     "mutation signIn($input: SignInInput!) {   signIn(input: $input) {     user {       name       __typename     }     __typename   } }",
				"variables": variable,
			})
			request, err = http.NewRequest("POST", url, bytes.NewBuffer(httpRequestBody))
			if err != nil {
				return "", err
			}
			request.Header.Set("Content-Type", "application/json")
			response, err = httpClient.Do(request)
			if err != nil {
				return "", err
			}
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
	fmt.Println(token)
	return token, nil
	// fmt.Println("Token:", Token)
	// os.Setenv("Token", Token)
	// time.Sleep(60 * time.Minute)
	// }
}
