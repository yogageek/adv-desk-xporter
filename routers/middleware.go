package routers

import (
	"fmt"
	"net/http"
	"os"
	"porter/db"
	gochan "porter/pkg/logic/gochan"
	vars "porter/pkg/logic/vars"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func middleware_api(c *gin.Context) {
	//查看是否正在做，如果是則值接返回錯誤
	if !vars.Get_PublicRes_State() {
		c.JSON(http.StatusLocked, gin.H{
			"error": "work is in process",
		})
		c.Abort() //當請求被中間件攔截時，確保停止後續的api調用
		return
	}

	fmt.Println("create log...")
	defer fmt.Println("saving log...")

	path := c.FullPath()
	pathList := strings.Split(path, "/")
	mode := pathList[len(pathList)-1]

	fileName := ""
	if mode == "import" {
		if file, err := c.Copy().FormFile("file"); err == nil {
			// d := fmt.Sprint(time.Now().UTC().Format("20060102"))
			// sysFileName := "iFactory_Desk_" + d
			fileName = file.Filename + ".json"
		}

	}
	if mode == "export" {
		d := fmt.Sprint(time.Now().UTC().Format("20060102"))
		sysFileName := "iFactory_Desk_" + d
		fileName = sysFileName + ".json"
	}

	// 2021/06/02 Username
	userName := "UNKNOWN"
	// Get username from request Token
	/*
		3. 身分認證(包含 Websocket)
		- 取得 API 請求後，取得請求中的 `IFPToken`，`EIToken` Cookies，然後帶上這些 Cookies 打 Desk Graphql API 取得該使用者的身分

		```graphql
			query me {
				me {
					user {
						id
						name
						username
					}
				}
			}
		```
	*/
	/*
	 IFPToken=eyXXX.eyXXX; EIToken=eyXXX.QZXXX
	*/
	ifpToken, _ := c.Cookie("IFPToken")
	// eiToken := c.Cookie("EIToken")
	// 這裡是直接取 Request 過來的 Cookie 去用 jwt 解 IFPToken, 並沒有做驗證!!並沒有做驗證!!並沒有做驗證!!
	// 後續還是要去打 Desk 的 Graphql 取得 me 作為 username!!

	token, _, err := new(jwt.Parser).ParseUnverified(ifpToken, jwt.MapClaims{})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		c.Abort() //當請求被中間件攔截時，確保停止後續的api調用
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["username"])
		userName = fmt.Sprint(claims["username"])
	} else {
		fmt.Println(err)
	}
	/*
		// Request Payload - graphql
		// [{"operationName":"me","variables":{},"query":"query me {\n  me {\n    user {\n      id\n      name\n      username\n      __typename\n    }\n    __typename\n  }\n}\n"}]
		// Response
		// [{"data":{"me":{"user":{"id":"USERID","name":"USERNAME","username":"user@domain","__typename":"User"},"__typename":"Me"}}}]
	*/
	// 2021/06/02 End

	// 2021/06/02 Username
	userName := "UNKNOWN"
	// Get username from request Token
	/*
		3. 身分認證(包含 Websocket)
		- 取得 API 請求後，取得請求中的 `IFPToken`，`EIToken` Cookies，然後帶上這些 Cookies 打 Desk Graphql API 取得該使用者的身分

		```graphql
			query me {
				me {
					user {
						id
						name
						username
					}
				}
			}
		```
	*/
	/*
	 IFPToken=eyXXX.eyXXX; EIToken=eyXXX.QZXXX
	*/
	ifpToken, _ := c.Cookie("IFPToken")
	// eiToken := c.Cookie("EIToken")
	// 這裡是直接取 Request 過來的 Cookie 去用 jwt 解 IFPToken, 並沒有做驗證!!並沒有做驗證!!並沒有做驗證!!
	// 後續還是要去打 Desk 的 Graphql 取得 me 作為 username!!

	token, _, err := new(jwt.Parser).ParseUnverified(ifpToken, jwt.MapClaims{})
	if err != nil {
		fmt.Println(err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["username"])
		userName = fmt.Sprint(claims["username"])
	} else {
		fmt.Println(err)
	}
	/*
		// Request Payload - graphql
		// [{"operationName":"me","variables":{},"query":"query me {\n  me {\n    user {\n      id\n      name\n      username\n      __typename\n    }\n    __typename\n  }\n}\n"}]
		// Response
		// [{"data":{"me":{"user":{"id":"USERID","name":"USERNAME","username":"user@domain","__typename":"User"},"__typename":"Me"}}}]
	*/
	// 2021/06/02 End

	log := &vars.Log{
		Database:  os.Getenv("MONGODB_DATABASE"),
		FileName:  fileName,
		Mode:      mode,
		CreatedAt: time.Now().UTC(),
		Result:    "success",
		UserName:  userName,
	}

	go ChanFlow()

	c.Next()

	// util.PrintJson(gochan.GetChan().Err)
	// util.PrintJson(gochan.GetChan().ErrCount)

	if gochan.GetChan().ErrCount > 0 {
		log.Result = "fail"
		log.Error = gochan.GetChan().Err
	}

	db.Insert(db.Clog, log)

}

func ChanFlow() {
	c := &gochan.LogChan{}
	c.InitChan()
	gochan.SetChan(c)              //設定全域chan
	go gochan.GetChan().TakeChan() //持續拿取
}

//havent implemet
func middleware_ws(c *gin.Context) {
	fmt.Println("exec middleware2...")
	defer fmt.Println("after exec middleware2...")
	c.Next()

}
