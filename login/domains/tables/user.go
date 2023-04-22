package tables

import "gorm.io/gorm"

type User struct {
	ID       int    `gorm:"column:id;primaryKey""`
	LoginID  string `gorm:"column:login_id"`
	Name     string `gorm:"column:name"`
	Password string `gorm:"column:password"`
	gorm.Model
}
