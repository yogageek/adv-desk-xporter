package routers

import (
	"fmt"
	"net/http"
	"os"
	"porter/db"
	gochan "porter/pkg/logic/gochan"
	vars "porter/pkg/logic/vars"
	util "porter/util"
	"strings"

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
			fileName = file.Filename
		}
	}
	if mode == "export" {
		fileName = "exportingData.json"
	}

	log := &vars.Log{
		Database:  os.Getenv("MONGODB_DATABASE"),
		FileName:  fileName,
		Mode:      mode,
		CreatedAt: util.GetNow(),
		Result:    "success",
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
