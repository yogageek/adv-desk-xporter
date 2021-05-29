package routers

import (
	"log"
	"net/http"

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
