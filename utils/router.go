package utils

import (
	"github.com/gin-gonic/gin"
	"go-basic/middleware"
	"go-basic/routes"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//Error middleware
	r.Use(middleware.ErrorMiddleware())
	//Cors middleware
	r.Use(middleware.CorsMiddleware)
	//Register web routes
	routes.RegisterWebRoutes(r)
	//You can register another route here

	return r
}
