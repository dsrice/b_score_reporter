package tables

import "gorm.io/gorm"

type Token struct {
	ID      int    `gorm:"column:id;primaryKey""`
	UserID  int    `gorm:"column:user_id"`
	User    User   `gorm:"foreignKey:UserID"`
	Auth    string `gorm:"column:auth"`
	Refresh string `gorm:"column:refresh"`
	gorm.Model
}
