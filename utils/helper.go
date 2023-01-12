package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// IsLocal 判断是否是本地开发环境
func IsLocal() bool {
	return os.Getenv("APP_ENV") == "local"
}

// IsDevelopment 判断是否是测试环境
func IsDevelopment() bool {
	return os.Getenv("APP_ENV") == "development"
}

// IsProduction 判断是否是生产环境
func IsProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}

// GetBasePath 获取根目录
func GetBasePath() string {
	//本地调试时用go run运行，用老方法获取root directory，其他编译运行的用新方法
	if IsLocal() {
		path, err := os.Getwd()
		if err != nil {
			logrus.Warn(err.Error())
			return ""
		}
		return path
	}
	ex, err := os.Executable()
	if err != nil {
		logrus.Warn(err.Error())
		return ""
	}
	exPath := filepath.Dir(ex)
	return exPath
}
