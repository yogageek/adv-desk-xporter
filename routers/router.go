package routers

import (
	v1 "porter/routers/api/porter/v1"

	"github.com/gin-gonic/gin"
)

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

		//use middlware
		api.Use(middleware_api)

		//api
		api.GET("/config/file/export", v1.Export)
		api.POST("/config/file/import", v1.Import)
		log.GET("/config/file/logs", v1.Logs)

		//ws
		ws.Use(middleware_ws)
		ws.GET("/config/file/status", v1.WsEvent).Use(middleware_ws)

		// 只能瀏覽 不能下載
		// StaticFile 是加載單個文件，而 StaticFS 是加載一個完整的目錄資源：
		// apiv1.StaticFile("/file/export", "./exportingData")
		// apiv1.StaticFS("/file/export", http.Dir(export.GetExcelFullPath()))
	}

	return r
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
