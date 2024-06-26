package bootstrap

import (
	"github.com/sirupsen/logrus"
	"go-basic/constants"
	"go-basic/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var gormLogger logger.Interface

func InitDB() {
	var driver = os.Getenv("DB_DRIVER")
	if utils.IsLocal() || utils.IsDevelopment() {
		gormLogger = logger.Default
	} else {
		gormLogger = logger.Discard
	}
	var err error
	logrus.Info("connecting db ....")
	var dialector gorm.Dialector
	switch driver {
	case "mysql":
		dialector = initMysql()
	case "sqlite":
		dialector = initSqlite()
	default:
		logrus.Fatal("Unsupported Driver: " + driver)
	}
	constants.DB, err = gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logrus.Fatalf("gorm connect error, error: %#v", err)
	}
	logrus.Info("connected")
}

func initMysql() gorm.Dialector {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") +
		":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?loc=Local&parseTime=true&charset=" +
		os.Getenv("DB_CHARSET") + "&collation=" + os.Getenv("DB_COLLATION")
	return mysql.Open(dsn)
}

func initSqlite() gorm.Dialector {
	dbPath := utils.GetBasePath() + "/assets/sqlite/sqlite.db"
	return sqlite.Open(dbPath)
}
