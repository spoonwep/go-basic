package middleware

import (
	"github.com/gin-gonic/gin"
	"go-basic/errors"
	"go-basic/response"
	auth "go-basic/service"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	token, err := GetToken(c)
	if err != nil {
		response.Fail(c, err)
		return
	}
	//获取并设置用户UID
	userID, _ := auth.GetUID(token)
	c.Set("UID", userID)
	c.Next()
}

func GetToken(c *gin.Context) (string, error) {
	//从headers中取出token
	authHeader := c.GetHeader("Authorization")
	prefix := "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", errors.WRONG_TOKEN_FORMAT
	}
	token := authHeader[strings.Index(authHeader, prefix)+len(prefix):]
	if token == "" {
		return "", errors.EMPTY_TOKEN
	} else {
		//验证token合法性
		_, err := auth.ParseToken(token)
		if err != nil {
			return "", errors.INVALID_TOKEN
		}
	}
	return token, nil
}
