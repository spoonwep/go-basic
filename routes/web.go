package routes

import (
	"github.com/gin-gonic/gin"
	"go-basic/controller"
)

func RegisterWebRoutes(r *gin.Engine) {
	r.GET("/", controller.IndexHandler)
}
