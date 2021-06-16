package routers

import (
	"log"
	"net/http"
	"os"
	"porter/config"
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

		//solve cors for 地端
		// http://127.0.0.1:11450/config/file/export
		hostname := c.Request.Host
		log.Println("Request Host with Port--->", hostname)
		ss := strings.Split(hostname, ":")
		host := ss[0]
		log.Println("Request Host--->", host)
		log.Println("Request Header Access-Control-Allow-Origin--->", c.GetHeader("Access-Control-Allow-Origin"))

		env_acao := os.Getenv("ACCESS_CONTROL_ALLOW_ORIGIN")
		log.Println("[ENV] ACCESS_CONTROL_ALLOW_ORIGIN--->", env_acao)
		if strings.Contains(env_acao, "HOSTNAME") {
			//地 ACCESS_CONTROL_ALLOW_ORIGIN會給http://${HOSTNAME}:10000
			acao := strings.ReplaceAll(env_acao, "${HOSTNAME}", host)
			log.Println("set new Access-Control-Allow-Origin--->", acao)
			c.Writer.Header().Set("Access-Control-Allow-Origin", acao)
		} else {
			//雲 ACCESS_CONTROL_ALLOW_ORIGIN會給https://ifp-organizer-${NAMESPACE}-${CLUSTER}.${DATACENTER}.wise-paas.com.cn
			acao := strings.ReplaceAll(env_acao, "${NAMESPACE}", config.Namespace)
			acao = strings.ReplaceAll(acao, "${CLUSTER}", config.Cluster)
			acao = strings.ReplaceAll(acao, "${DATACENTER}", config.Datacenter)
			log.Println("set new Access-Control-Allow-Origin--->", acao)
			c.Writer.Header().Set("Access-Control-Allow-Origin", acao)
		}

		// defer func() {
		// 	if err := recover(); err != nil {
		// 		log.Printf("Panic info is: %v", err)
		// 	}
		// }()

		c.Next()
	}
}
