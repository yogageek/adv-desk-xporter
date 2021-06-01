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
	// Token  string
	// Token2 string

	//for debugging
	Token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJpZnAub3JnIiwic3ViIjoiNjAzNWVjZDk3OTA4ZTUwMDA3YTMxZGU4IiwiYXVkIjoidXNlciIsInVzZXJuYW1lIjoiZGV2YW5saWFuZ0BpaWkub3JnLnR3Iiwid2lzZVBhYXNSZWZyZXNoVG9rZW4iOiJlNzI0MzAzOC1jMmI5LTExZWItOTU5NS00MmQ4MmE1MWExMzQiLCJpYXQiOjE2MjI1Mzg5MjYsImV4cCI6MTYyMjU2NzcyNn0.vFgdS2MxOjPm4VgVgh-YCgadMZ4RASk3vHAZcwDZiTQ;EIToken=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJjb3VudHJ5IjoiVFciLCJjcmVhdGlvblRpbWUiOjE1ODQ5MzQ5MzksImV4cCI6MTYyMjU0MjUyNiwiZmlyc3ROYW1lIjoiRGV2YW4iLCJpYXQiOjE2MjI1Mzg5MjYsImlkIjoiNGMzNGM2MzgtNmNiOC0xMWVhLWIzOGUtYmFjYmEwYzcyYzczIiwiaXNzIjoid2lzZS1wYWFzIiwibGFzdE1vZGlmaWVkVGltZSI6MTYyMjUzODkwMywibGFzdE5hbWUiOiJMaWFuZyIsInJlZnJlc2hUb2tlbiI6ImU3MjQzMDM4LWMyYjktMTFlYi05NTk1LTQyZDgyYTUxYTEzNCIsInN0YXR1cyI6IkFjdGl2ZSIsInVzZXJuYW1lIjoiZGV2YW5saWFuZ0BpaWkub3JnLnR3In0.qiV51B8jRBP4z9oDGo4Q1ItYai_1xdr7dOROD3hBwXgp3Sf48D8ZjS9ZZqSUNOPiqkZboZvaONUeQH7BQjYLBw"
	Token2 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJpZnAub3JnIiwic3ViIjoiNjAzNWVjZDk3OTA4ZTUwMDA3YTMxZGU4IiwiYXVkIjoidXNlciIsInVzZXJuYW1lIjoiZGV2YW5saWFuZ0BpaWkub3JnLnR3Iiwid2lzZVBhYXNSZWZyZXNoVG9rZW4iOiJlNzI0MzAzOC1jMmI5LTExZWItOTU5NS00MmQ4MmE1MWExMzQiLCJpYXQiOjE2MjI1Mzg5MjYsImV4cCI6MTYyMjU2NzcyNn0.vFgdS2MxOjPm4VgVgh-YCgadMZ4RASk3vHAZcwDZiTQ;EIToken=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJjb3VudHJ5IjoiVFciLCJjcmVhdGlvblRpbWUiOjE1ODQ5MzQ5MzksImV4cCI6MTYyMjU0MjUyNiwiZmlyc3ROYW1lIjoiRGV2YW4iLCJpYXQiOjE2MjI1Mzg5MjYsImlkIjoiNGMzNGM2MzgtNmNiOC0xMWVhLWIzOGUtYmFjYmEwYzcyYzczIiwiaXNzIjoid2lzZS1wYWFzIiwibGFzdE1vZGlmaWVkVGltZSI6MTYyMjUzODkwMywibGFzdE5hbWUiOiJMaWFuZyIsInJlZnJlc2hUb2tlbiI6ImU3MjQzMDM4LWMyYjktMTFlYi05NTk1LTQyZDgyYTUxYTEzNCIsInN0YXR1cyI6IkFjdGl2ZSIsInVzZXJuYW1lIjoiZGV2YW5saWFuZ0BpaWkub3JnLnR3In0.qiV51B8jRBP4z9oDGo4Q1ItYai_1xdr7dOROD3hBwXgp3Sf48D8ZjS9ZZqSUNOPiqkZboZvaONUeQH7BQjYLBw"
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

wait:
	if Token == "" && Token2 == "" {
		time.Sleep(time.Millisecond * 300)
	} else {
		return
	}
	goto wait

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
