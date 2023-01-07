package controller

import (
	"github.com/gin-gonic/gin"
	"go-basic/response"
)

func IndexHandler(c *gin.Context) {
	response.Success(c, "Hello World")
}
