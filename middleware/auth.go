package middleware

import (
	"github.com/gin-gonic/gin"
	"go-basic/errors"
	"go-basic/response"
	auth "go-basic/service"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	//从headers中取出token
	authHeader := c.GetHeader("Authorization")
	prefix := "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		response.Fail(c, errors.WRONG_TOKEN_FORMAT)
		return
	}
	token := authHeader[strings.Index(authHeader, prefix)+len(prefix):]
	if token == "" {
		response.Fail(c, errors.EMPTY_TOKEN)
		return
	} else {
		//验证token合法性
		_, err := auth.ParseToken(token)
		if err != nil {
			response.Fail(c, errors.INVALID_TOKEN)
			return
		}
	}
	//获取并设置用户UID
	userID, _ := auth.GetUID(token)
	c.Set("UID", userID)
	c.Next()
}
