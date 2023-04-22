package database

import (
	mysqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"login/domains/tables"
	"os"
	"time"
)

func NewDataBase() *gorm.DB {
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
	db.AutoMigrate(&tables.User{})
	db.AutoMigrate(&tables.Token{})
	db.AutoMigrate(&tables.Login{})
}
