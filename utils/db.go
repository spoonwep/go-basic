package utils

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") +
		":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?loc=Local&parseTime=true&charset=" +
		os.Getenv("DB_CHARSET") + "&collation=" + os.Getenv("DB_COLLATION")
	logrus.Info("connecting db ....")

	var gormLogger logger.Interface
	if IsLocal() || IsDevelopment() {
		gormLogger = logger.Default
	} else {
		gormLogger = logger.Discard
	}
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logrus.Fatalf("gorm connect error, error: %#v", err)
	}
	logrus.Info("connected")
}
