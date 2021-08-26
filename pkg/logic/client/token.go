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
	"porter/util"
	"strings"
	"time"

	"github.com/beevik/ntp"
	"github.com/bitly/go-simplejson"
	"github.com/prometheus/common/log"
)

var (
	TaipeiTimeZone, _ = time.LoadLocation("Asia/Taipei")
)

var (
	UserPwdToken  string
	UserPwdToken2 string
)

func Loop_RefreshTokenByUserPwd() {

	go func() {
		for {
			fmt.Println("<<< Loop refresh token >>>")
			var err error
			UserPwdToken, err = RefreshTokenByUserPwd(config.IFP_URL)
			if err != nil {
				log.Error("RefreshToken IFP_URL fail")
				panic(err)
			}
			UserPwdToken2, err = RefreshTokenByUserPwd(config.IFP_URL_IN)
			if err != nil {
				log.Error("RefreshToken IFP_URL_IN fail")
				panic(err)
			}
			time.Sleep(15 * time.Minute)
		}
	}()

wait:
	if UserPwdToken == "" && UserPwdToken2 == "" {
		time.Sleep(time.Millisecond * 300)
	} else {
		return
	}
	goto wait
}

func RefreshTokenByUserPwd(url string) (token string, err error) {
	httpClient := &http.Client{}
	content := map[string]string{"username": config.IfpDeskUsername, "password": config.IfpDeskPassword}
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
	errors := m.Get("errors")
	errorsb, _ := errors.MarshalJSON()
	if string(errorsb) != "null" {
		log.Error("graphql sign in response got errors...")
		panic(string(errorsb))
	}

	// r, _ := m.MarshalJSON()
	// fmt.Printf("response body: %s", r)

	header := response.Header
	// fmt.Printf("response header: %s", header)
	cookie := header["Set-Cookie"]
	tempSplit := strings.Split(cookie[0], ";")
	ifpToken := tempSplit[0]
	tempSplit = strings.Split(cookie[1], ";")
	eiToken := tempSplit[0]
	token = ifpToken + ";" + eiToken
	fmt.Println("Token from cookie:", token)
	return token, nil
}

// 2021/05/31 Add AppSecret Token
func RefreshTokenByAppSecret() {
	for {
		fmt.Println("----------", time.Now().In(config.TaipeiTimeZone), "----------")
		if len(config.AppSecretFile) > 0 {
			util.PrintBlue("running on an enviroment which is not cloud-------------")
			appSecret, err := os.ReadFile(config.AppSecretFile)
			if err != nil {
				fmt.Println("Read appSecret_file error:", err)
			} else {
				config.Token = strings.Trim(string(appSecret), "\n")
			}
			fmt.Println("Token from appSecret:", config.Token)
			time.Sleep(60 * time.Minute)
		} else {
			fmt.Println("running on an enviroment which is cloud--------------")
			//timestamp := time.Now()
			//options := &newSRPTokenOptions{Timestamp: &timestamp}
			result := newSrpToken(config.ServiceName, nil)
			httpClient := &http.Client{}
			request, _ := http.NewRequest("GET", config.SSOURL+"/clients/"+config.ClientName, nil)
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("X-Auth-SRPToken", result)
			fmt.Println("Set X-Auth-SRPToken: ", result)
			q := request.URL.Query()
			q.Add("cluster", config.Cluster)
			q.Add("workspace", config.Workspace)
			q.Add("namespace", config.Namespace)
			q.Add("appId", config.AppID)
			q.Add("serviceName", config.ServiceName)
			request.URL.RawQuery = q.Encode()
			response, err := httpClient.Do(request)
			if err != nil {
				panic(err)
			}

			fmt.Printf("request url and param: %s \n", request.URL)
			fmt.Printf("response: %+v \n", response)
			m, err := simplejson.NewFromReader(response.Body)
			if err != nil {
				panic(err)
			}
			config.Token = m.Get("clientSecret").MustString()
			fmt.Println("Token from SSO:", config.Token)
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

/*
● base64Url_encode(AES_encode(timestamp-srpName)) (tmestamp為當前10位時間戳記)

● AES_encode
	■ AES/ECB/PKCS5Padding
○ Secret key: ssoisno12345678987654321
○ Golden unit test
	■ Source of plain text: 1234567890-SCADA
	■ base64Encode(AES encode(src)): (可利用 Online Crypto Tool 驗證)
		● h9mmu4CIc+YwBWDamtMKMA9tdDQNzz/RLTFfsfGoQhg=
	■ base64UrlEncode(AES encode(src)): (實際帶在 header裡請用 base64UrlEncode)
		● h9mmu4CIc-YwBWDamtMKMA9tdDQNzz_RLTFfsfGoQhg= ○ Online Crypto Tool (base64Encode)
	■ https://goo.gl/6ig1No ○ Java Sample code
*/
func newSrpToken(serviceName string, opts ...*newSRPTokenOptions) string {
	ntpTime, err := ntp.Time("time.google.com")
	if err != nil {
		fmt.Println(err)
	}
	now := ntpTime

	timestamp := &now

	key := "ssoisno12345678987654321"
	fmt.Println(timestamp.Unix())
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
