package utils

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go-basic/middleware"
	"go-basic/routes"
)

func InitRouter() *gin.Engine {
	if IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	//pprof middleware
	if !IsProduction() {
		pprof.Register(r)
	}
	//Error middleware
	r.Use(middleware.ErrorMiddleware())
	//Cors middleware
	r.Use(middleware.CorsMiddleware)
	//Register web routes
	routes.RegisterWebRoutes(r)
	//You can register another route here

	return r
}
