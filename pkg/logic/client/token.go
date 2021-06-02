package logic

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"porter/config"
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

// 2021/05/31 Add AppSecret Token
func RefreshTokenByAppSecret() {
	for {
		fmt.Println("----------", time.Now().In(config.TaipeiTimeZone), "----------")
		if len(config.Datacenter) == 0 {
			fmt.Println("refreshToken")
			httpClient := &http.Client{}
			content := map[string]string{"username": config.AdminUsername, "password": config.AdminPassword}
			variable := map[string]interface{}{"input": content}
			httpRequestBody, _ := json.Marshal(map[string]interface{}{
				"query":     "mutation signIn($input: SignInInput!) {   signIn(input: $input) {     user {       name       __typename     }     __typename   } }",
				"variables": variable,
			})
			request, _ := http.NewRequest("POST", config.IFPURL, bytes.NewBuffer(httpRequestBody))
			request.Header.Set("Content-Type", "application/json")
			response, _ := httpClient.Do(request)
			m, _ := simplejson.NewFromReader(response.Body)
			for {
				if len(m.Get("errors").MustArray()) == 0 {
					break
				} else {
					fmt.Println("----------", time.Now().In(config.TaipeiTimeZone), "----------")
					fmt.Println("retry refreshToken")
					httpRequestBody, _ = json.Marshal(map[string]interface{}{
						"query":     "mutation signIn($input: SignInInput!) {   signIn(input: $input) {     user {       name       __typename     }     __typename   } }",
						"variables": variable,
					})
					request, _ = http.NewRequest("POST", config.IFPURL, bytes.NewBuffer(httpRequestBody))
					request.Header.Set("Content-Type", "application/json")
					response, _ = httpClient.Do(request)
					m, _ = simplejson.NewFromReader(response.Body)
					time.Sleep(6 * time.Minute)
				}
			}
			header := response.Header
			cookie := header["Set-Cookie"]
			var ifpToken, eiToken string
			for _, cookieContent := range cookie {
				tempSplit := strings.Split(cookieContent, ";")
				if strings.HasPrefix(tempSplit[0], "IFPToken") {
					ifpToken = tempSplit[0]
				} else if strings.HasPrefix(tempSplit[0], "EIToken") {
					eiToken = tempSplit[0]
				}
			}
			if eiToken == "" {
				config.Token = ifpToken
			} else {
				config.Token = ifpToken + ";" + eiToken
			}
			fmt.Println("Token:", config.Token)
			time.Sleep(60 * time.Minute)
		} else {
			/*
				- 呼叫 Get EnSaaS Client API，失敗就等十秒後重試，三次後就自殺
				- https://docs.wise-paas.advantech.com.cn/zh-tw/Guides_and_API_References/Cloud_Services/SSO/1569378270407013482/v1.4.6  中的 GET /clients/{clientIdOrName}

				- HTTP Method: GET
				- HTTP Header: X-Auth-SRPToken=${WISE_PAAS_SERVICE_NAME}
				- URL: ${WISE_PAAS_SSO_API_URL}/clients/${WISE_PAAS_SERVICE_NAME}
				- Query String
				- mode=name
				- namespace=${namespace}
				- serviceName=${WISE_PAAS_SERVICE_NAME}
				- appId=${appID}
				- datacenter=${datacenter}
				- workspace=${workspace}
				- cluster=${cluster}

				- 取回應中的 clientSecret 欄位

				```json
				{
					"clientId": "string",
					"clientSecret": "string",
					"creationTime": 0,
					"lastModifiedTime": 0,
					"appName": "string",
					"appId": "string",
					"serviceName": "string",
					"cluster": "string",
					"workspace": "string",
					"namespace": "string",
					"datacenter": "string",
					"redirectUrl": "string",
					"scopes": [
					"string"
					]
				}
				```
			*/
			fmt.Println("refreshClientSecret")
			timestamp := time.Now()
			options := &newSRPTokenOptions{Timestamp: &timestamp}
			result := newSrpToken("OEE", options)
			httpClient := &http.Client{}
			request, _ := http.NewRequest("GET", config.SSOURL+"/clients/OEE", nil)
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("X-Auth-SRPToken", result)
			q := request.URL.Query()
			if config.Namespace == "ifpsdev" || config.Namespace == "ifpsdemo" {
				// 我們自己環境連的是 eks011 training 的站點, 這裡寫死
				q.Add("cluster", "eks011")
				q.Add("workspace", "53e8c8bd-b724-4c87-a905-5bbc5c30a36c")
				q.Add("namespace", "training")
			} else {
				q.Add("cluster", config.Cluster)
				q.Add("workspace", config.Workspace)
				q.Add("namespace", config.Namespace)
			}
			q.Add("serviceName", "OEE")
			request.URL.RawQuery = q.Encode()
			response, _ := httpClient.Do(request)
			m, _ := simplejson.NewFromReader(response.Body)
			config.Token = m.Get("clientSecret").MustString()
			fmt.Println("Token:", config.Token)
			time.Sleep(60 * time.Minute)
		}
	}
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

type ecbEncrypter ecb

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

func newECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

type newSRPTokenOptions struct {
	Timestamp *time.Time
}

// PKCS7Padding adds padding to data
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func newSrpToken(serviceName string, opts ...*newSRPTokenOptions) string {
	now := time.Now()
	timestamp := &now

	for _, opt := range opts {
		if opt.Timestamp != nil {
			timestamp = opt.Timestamp
		}
	}

	key := "ssoisno12345678987654321"
	src := fmt.Sprintf("%v-%v", timestamp.Unix(), serviceName)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	data := PKCS7Padding([]byte(src), blockSize)

	encryptData := make([]byte, len(data))

	ecb := newECBEncrypter(block)
	ecb.CryptBlocks(encryptData, data)

	token := base64.URLEncoding.EncodeToString(encryptData)
	token = strings.ReplaceAll(token, "=", "")
	return token
}
