package middleware

import (
	"github.com/gin-gonic/gin"
	"go-basic/errors"
	auth "go-basic/service"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	//从headers中取出token
	response := errors.NORMAL
	token := c.GetHeader("token")
	if token == "" {
		response = errors.EMPTY_TOKEN
	} else {
		//验证token合法性
		_, err := auth.ParseToken(token)
		if err != nil {
			response = errors.INVALID_TOKEN
		}
	}
	if response.Code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code": response.Code,
			"msg":  response.Msg,
			"data": nil,
		})
		c.Abort()
		return
	}
	//霍去病设置用户UID
	userID, _ := auth.GetUID(token)
	c.Set("UID", userID)
	c.Next()
}
