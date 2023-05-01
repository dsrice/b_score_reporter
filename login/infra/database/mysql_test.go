package database_test

import (
	"github.com/stretchr/testify/assert"
	"login/infra/database"
	"login/infra/env"
	"testing"
)

func TestNewDataBase(t *testing.T) {
	env.LoadEnv()

	db := database.NewDataBase()

	assert.NotNil(t, db)
}
