package env_test

import (
	"testing"

	"github.com/Linkinlog/loggr/internal/env"
	"github.com/stretchr/testify/assert"
)

func TestNewEnv(t *testing.T) {
	t.Parallel()

	env := env.NewEnv()
	assert.NotNil(t, env)
}
