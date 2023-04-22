package env_test

import (
	"github.com/stretchr/testify/assert"
	"login/infra/env"
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	env.LoadEnv()

	result := os.Getenv("DBNAME")
	assert.Equal(t, result, "b_score")
}
