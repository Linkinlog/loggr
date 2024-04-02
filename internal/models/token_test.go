package models_test

import (
	"testing"
	"time"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	t.Parallel()
	tooShort := models.NewToken(0, 14*24*time.Hour)

	assert.Equal(t, models.MaxTTL, tooShort.TTL)
	assert.NotEmpty(t, tooShort.Value)
	assert.Len(t, tooShort.Value, models.MinTokenLength*2)
}

func TestToken_String(t *testing.T) {
	t.Parallel()
	token := models.NewToken(32, 14*24*time.Hour)

	assert.Equal(t, token.Value, token.String())
}

func TestToken_Expired(t *testing.T) {
	t.Parallel()
	token := models.NewToken(32, 14*24*time.Hour)

	assert.False(t, token.Expired())
}
