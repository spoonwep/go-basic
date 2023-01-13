package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-basic/constants"
	"go-basic/errors"
	auth "go-basic/service"
	"net/http"
	"strconv"
)

type CasbinMiddleware struct {
	enforcer *casbin.Enforcer
}

// EnsurePermissions 检查权限
func (md *CasbinMiddleware) EnsurePermissions(c *gin.Context) (bool, string, error) {
	token, err := GetToken(c)
	if err != nil {
		return false, "", err
	}
	userId, _ := auth.GetUID(token)
	method := c.Request.Method
	path := c.Request.URL.Path
	result, err := md.enforcer.Enforce(strconv.Itoa(int(userId)), path, method)
	if err != nil {
		logrus.Warn(err.Error())
		return false, "", nil
	}
	return result, token, nil
}

func NewCasbinMiddleware() gin.HandlerFunc {
	e, err := casbin.NewEnforcer(constants.BasePath+"/assets/casbin/model.conf", constants.BasePath+"/assets/casbin/policy.csv")
	if err != nil {
		logrus.Fatal(err.Error())
		return nil
	}
	md := &CasbinMiddleware{e}
	return func(c *gin.Context) {
		//检查权限
		passed, token, err := md.EnsurePermissions(c)
		if err != nil {
			_ = c.AbortWithError(200, err)
			return
		}
		if !passed {
			_ = c.AbortWithError(http.StatusForbidden, errors.NO_PERMISSION)
		}
		//获取并设置用户UID
		userID, _ := auth.GetUID(token)
		c.Set("UID", userID)
	}
}
