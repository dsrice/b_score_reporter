package database

import (
	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"login/domains/tables"
	"login/infra/env"
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

	println(c.FormatDSN())
	db, err := gorm.Open(mysql.Open(c.FormatDSN()), &gorm.Config{})

	if err != nil {
		println(err.Error())
		return nil
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
