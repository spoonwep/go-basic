package response

import (
	"github.com/gin-gonic/gin"
	"go-basic/errors"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": errors.NORMAL.Code,
		"msg":  errors.NORMAL.Msg,
		"data": data,
	})
}

func Fail(c *gin.Context, err error) {
	_ = c.Error(err)
}
