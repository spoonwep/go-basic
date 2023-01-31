package controller

import (
	"github.com/gin-gonic/gin"
	"go-basic/delivery/response"
)

func IndexHandler(c *gin.Context) {
	response.Success(c, "Hello World")
}
