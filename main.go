package main

import (
	"github.com/sirupsen/logrus"
	"go-basic/bootstrap"
	"go-basic/constants"
	"go-basic/utils"
	"os"
)

func main() {
	//初始化配置文件
	bootstrap.InitConfig()
	//初始化logger
	bootstrap.InitLogger()
	//初始化Database
	bootstrap.InitDB()
	//初始化Redis
	if os.Getenv("REDIS_HOST") != "" {
		bootstrap.InitRedis()
	}
	//初始化定时任务（需要在web服务之前初始化，否则web服务无法启动）
	bootstrap.InitCronjob()
	//初始化根目录
	initRoot()
	//开启WebServer
	startWebServer()
}

func startWebServer() {
	//初始化路由
	router := bootstrap.InitRouter()
	//初始化validator
	bootstrap.InitValidator()
	//启动WebServer
	var err error
	if os.Getenv("ENABLE_HTTPS") == "true" {
		err = router.RunTLS(":"+os.Getenv("APP_TLS_PORT"), os.Getenv("CERT_PATH"), os.Getenv("KEY_PATH"))
	} else {
		err = router.Run(":" + os.Getenv("APP_HTTP_PORT"))
	}
	if err != nil {
		logrus.Fatal(err.Error())
	}
}

func initRoot() {
	constants.BasePath = utils.GetBasePath()
}
