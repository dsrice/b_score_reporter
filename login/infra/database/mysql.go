package database

import (
	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"login/domains/tables"
	"login/infra/env"
	zaplogger "login/infra/logger"
	"moul.io/zapgorm2"
	"os"
	"time"
)

func NewDataBase() *gorm.DB {
	env.LoadEnv()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		println(err.Error())
		return nil
	}

	c := mysqldriver.Config{
		DBName:               os.Getenv("DBNAME"),
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASSWORD"),
		Addr:                 os.Getenv("DBADDR"),
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}

	zLogger := zapgorm2.New(zaplogger.NewLogger())

	db, err := gorm.Open(mysql.Open(c.FormatDSN()), &gorm.Config{
		Logger: zLogger,
	})
	if err != nil {
		println(err.Error())
		return nil
	}

	db.Logger = db.Logger.LogMode(logger.Info)
	err = SetCallBack(db)

	if err != nil {
		println(err.Error())
	}

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	execMigrate(db, &tables.User{})
	execMigrate(db, &tables.Token{})
	execMigrate(db, &tables.Login{})
}

func execMigrate(db *gorm.DB, target interface{}) {
	err := db.AutoMigrate(target)

	if err != nil {
		log.Error(err)
	}
}
