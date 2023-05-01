package logger_test

import (
	"github.com/stretchr/testify/assert"
	"login/infra/logger"
	"testing"
)

func TestNewLogger(t *testing.T) {
	log := logger.NewLogger()

	assert.NotNil(t, log)
}
