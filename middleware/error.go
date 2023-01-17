package middleware

import (
	"github.com/gin-gonic/gin"
	"go-basic/errors"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		//此处要执行完所有任务后检查context中error有没有
		for _, e := range c.Errors {
			err := e.Err
			//区分自定义错误和非自定义错误
			if baseErr, ok := err.(*errors.BaseError); ok {
				c.JSON(http.StatusOK, gin.H{
					"code": baseErr.Code,
					"msg":  baseErr.Msg,
					"data": baseErr.Data,
				})
				return
			} else {
				headerCode := c.Writer.Status()
				c.JSON(headerCode, gin.H{
					"code": errors.SERVER_ERROR.Code,
					"msg":  errors.SERVER_ERROR.Msg,
					"data": err.Error(),
				})
				return
			}
		}
	}
}
