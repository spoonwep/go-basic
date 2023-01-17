package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// AddPolicy 添加并保存规则
func AddPolicy(c *gin.Context, params ...interface{}) error {
	enforcer, _ := c.Get("enforcer")
	e := enforcer.(*casbin.Enforcer)
	if ok, err := e.AddPolicy(params...); !ok {
		return err
	}
	err := e.SavePolicy()
	if err != nil {
		return err
	}
	return nil
}

// RemovePolicy 删除一条规则
func RemovePolicy(c *gin.Context, params ...interface{}) error {
	enforcer, _ := c.Get("enforcer")
	e := enforcer.(*casbin.Enforcer)
	if ok, err := e.RemovePolicy(params...); !ok {
		return err
	}
	err := e.SavePolicy()
	if err != nil {
		return err
	}
	return nil
}
