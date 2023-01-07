package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLogger() {
	logrus.SetOutput(os.Stdout)
	//显示所在文件
	logrus.SetReportCaller(true)
	//生产环境用json输出，可兼容logstash
	if IsProduction() {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}
