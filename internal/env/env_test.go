package env_test

import (
	"os"
	"testing"

	"github.com/Linkinlog/loggr/internal/env"
	"github.com/stretchr/testify/assert"
)

func TestNewEnv(t *testing.T) {
	t.Parallel()

	env := env.NewEnv()
	assert.NotNil(t, env)
}

func TestGet(t *testing.T) {
	t.Parallel()

	env := env.NewEnv()
	assert.NotNil(t, env)

	key := "TEST"
	value := "test"
	os.Setenv(key, value)
	defer os.Unsetenv(key)

	assert.Equal(t, value, env.GetOrDefault(key, "wrong"))
	assert.Equal(t, "right", env.GetOrDefault("wrong", "right"))
}
