package bootstrap

import (
	"github.com/sirupsen/logrus"
	"go-basic/utils"
	"os"
)

func InitLogger() {
	logrus.SetOutput(os.Stdout)
	//显示所在文件
	logrus.SetReportCaller(true)
	//生产环境用json输出，可兼容logstash
	if utils.IsProduction() {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}
