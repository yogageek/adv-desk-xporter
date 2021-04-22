package v1

import (
	"net/http"
	"porter/db"
	vars "porter/pkg/logic/vars"

	"github.com/gin-gonic/gin"
)

// GET /file/status
func Logs(c *gin.Context) {
	var i []vars.Log
	if err := db.FindAll(db.Clog, nil, nil, &i); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, i)
}
