package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"porter/db"
	vars "porter/pkg/logic/vars"
	v1 "porter/routers/api/porter/v1"
	util "porter/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// origin := c.Request.Header.Get("Origin") //請求頭部
		// if origin != "" {
		//接收客戶端傳送的origin （重要！）
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//伺服器支援的所有跨域請求的方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		//允許跨域設定可以返回其他子段，可以自定義欄位
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session")
		// 允許瀏覽器（客戶端）可以解析的頭部 （重要）
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		//設定快取時間
		c.Header("Access-Control-Max-Age", "172800")
		//允許客戶端傳遞校驗資訊比如 cookie (重要)
		c.Header("Access-Control-Allow-Credentials", "true")
		// }

		//允許型別校驗
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "method == OPTIONS ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(Cors())

	// gin.SetMode(setting.RunMode)

	//----------------->
	api := r.Group("/")
	log := r.Group("/")
	ws := r.Group("/")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})
		api.Use(middlewareAPI) //新增存log功能 //.use的先後放置順序很重要 會影響以下
		api.GET("/config/file/export", v1.Export)
		api.POST("/config/file/import", v1.Import)

		log.GET("/config/file/logs", v1.Logs)

		//ws
		ws.Use(middlewareWS) //新增存log功能 //.use的先後放置順序很重要 會影響以下
		ws.GET("/config/file/status", v1.WsEvent).Use(middlewareWS)
		// apiv1.GET("/file/status", v1.Status) //move to websocket

		// 只能瀏覽 不能下載
		// StaticFile 是加載單個文件，而 StaticFS 是加載一個完整的目錄資源：
		// apiv1.StaticFile("/file/export", "./exportingData")
		// apiv1.StaticFS("/file/export", http.Dir(export.GetExcelFullPath()))
	}

	return r
}

func middlewareAPI(c *gin.Context) {
	fmt.Println("create log...")
	defer fmt.Println("saving log...")

	path := c.FullPath()
	pathList := strings.Split(path, "/")
	mode := pathList[len(pathList)-1]

	fileName := ""
	if mode == "import" {
		if file, err := c.Copy().FormFile("file"); err == nil {
			fileName = file.Filename
		}
	}

	log := &vars.Log{
		Database:  os.Getenv("MONGODB_DATABASE"),
		FileName:  fileName,
		Mode:      mode,
		CreatedAt: util.GetNow(),
	}

	c.Next()

	log.Result = "success"
	log.Error = ""
	//取error出來

	db.Insert(db.Clog, log)

}

func middlewareWS(c *gin.Context) {
	fmt.Println("exec middleware2...")
	defer fmt.Println("after exec middleware2...")
	c.Next()
}

/*

   3. 讀取列表 API 要做分頁，以及 filter 狀態(成功、執行中、失敗)和種類(匯入匯出)
   4. 匯入匯出 Collection 欄位建議(看你有沒有其他建議或缺失)
      1. 資料庫 ID
      2. filename (你先自己定義格式)
      3. username (建立工作時會傳給你)
      4. 建立時間 createdAt
      5. 種類: import or export
      6. 狀態: 成功、執行中、失敗
      7. error: 整個 Error 物件，除了以下階段以外丟出的錯誤，e.g. 檢查語言失敗
      8. machineStatus
         1. loaded
         2. total
         3. error 整個 Error 物件，Machine Status 階段丟出的錯誤
      9. mappingRule
         1. loaded
         2. total
         3. error 整個 Error 物件，Mapping Rule 階段丟出的錯誤
      10. ...其他階段
*/
