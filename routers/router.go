package routers

import (
	v1 "porter/routers/api/porter/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	// gin.SetMode(setting.RunMode)

	//----------------->
	apiv1 := r.Group("/")
	{
		apiv1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})
		apiv1.GET("/config/file/export", v1.Export)
		apiv1.POST("/config/file/import", v1.Import)
		// apiv1.GET("/file/status", v1.Status) //move to websocket

		// 只能瀏覽 不能下載
		// StaticFile 是加載單個文件，而 StaticFS 是加載一個完整的目錄資源：
		// apiv1.StaticFile("/file/export", "./exportingData")
		// apiv1.StaticFS("/file/export", http.Dir(export.GetExcelFullPath()))
	}

	return r
}
