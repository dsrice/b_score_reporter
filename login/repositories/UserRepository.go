package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
	"login/domains/tables"
	"login/infra/jaeger"
	"login/repositories/adapters"
)

type userRepository struct {
	c  echo.Context
	db *gorm.DB
}

func NewUser(db *gorm.DB) adapters.UserAdapter {
	return &userRepository{
		db: db,
	}
}

// Create
// ユーザー作成処理
func (repo *userRepository) Create(user tables.User) error {
	err := repo.db.Create(&user)

	if err.Error != nil {
		return err.Error
	}

	return nil
}

// GetUser
// ユーザー取得処理
func (repo *userRepository) GetUser(sc tables.User) ([]*tables.User, error) {
	sp := jaeger.CreateChileSpan(repo.c, "Repository")
	defer sp.Finish()

	var userList []*tables.User

	err := repo.db.Where(&sc).Find(&userList)

	sp.LogFields(log.String("status", "start"))
	if err.Error != nil {
		return nil, err.Error
	}

	return userList, nil
}

func (repo *userRepository) SetContext(c echo.Context) {
	repo.c = c
	repo.db.Statement.Context = c.Request().Context()
}
