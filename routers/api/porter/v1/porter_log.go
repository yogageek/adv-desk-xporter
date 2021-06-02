package v1

import (
	"net/http"
	"porter/db"
	vars "porter/pkg/logic/vars"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /file/status
func Logs(c *gin.Context) {
	var i []vars.Log
	// 2021/06/01 Query param
	/*
			2. Logs API 參數
		   - 新增 limit 參數
		   - 新增 sort 參數
		   - 新增 username 欄位(做完第3項身分驗證，就可以知道 username)
	*/
	limit, _ := strconv.Atoi(c.Query("limit"))
	sort := c.Query("sort")
	if sort == "" {
		sort = "-_id"
	}
	skip, _ := strconv.Atoi(c.Query("page"))
	// queryM := bson.M{}
	// for key, val := range c.Request.URL.Query() {
	// 	if !util.InArray(key, []string{"_limit", "_offset", "_sort"}) {
	// 		queryM[key] = val
	// 	}

	// }
	if err := db.FindAllByOpts(db.Clog, nil, nil, &i, sort, skip, limit); err != nil {
		// 2021/06/01 End
		//if err := db.FindAll(db.Clog, nil, nil, &i); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, i)
}
