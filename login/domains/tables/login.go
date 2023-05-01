package tables

import (
	"gorm.io/gorm"
	"time"
)

type Login struct {
	ID            int       `gorm:"column:id;primaryKey"`
	UserID        int       `gorm:"column:user_id"`
	User          User      `gorm:"foreignKey:UserID"`
	ErrorCount    int       `gorm:"column:error_count"`
	LastLoginDate time.Time `gorm:"column:last_login_date"`
	gorm.Model
}
