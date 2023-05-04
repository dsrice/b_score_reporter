package test

import (
	"gorm.io/gorm"
	"login/infra/database"
)

var DB *gorm.DB = nil

func DataBase() *gorm.DB {
	if DB == nil {
		DB = database.NewDataBase()
	}

	return DB
}
