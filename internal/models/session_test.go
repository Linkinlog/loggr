package models_test

import (
	"testing"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestNewSession(t *testing.T) {
	t.Parallel()
	u, _ := models.NewUser("Batman", "batman@hotmail.com", "password123", "")
	s := models.NewSession(u)

	assert.NotNil(t, s)
	assert.Equal(t, u, s.User)
}

func TestSession_Id(t *testing.T) {
	t.Parallel()
	u, _ := models.NewUser("Batman", "batman@hotmail.com", "password123", "")
	s := models.NewSession(u)

	assert.NotEmpty(t, s.Id)
}

func TestSession_TTL(t *testing.T) {
	t.Parallel()
	u, _ := models.NewUser("Batman", "batman@hotmail.com", "password123", "")
	s := models.NewSession(u)

	assert.NotZero(t, s.TTL())
}

func TestSession_ToCookie(t *testing.T) {
	t.Parallel()
	u, _ := models.NewUser("Batman", "batman@hotmail.com", "password123", "")
	s := models.NewSession(u)

	cookie := s.ToCookie()

	assert.NotNil(t, cookie)
	assert.Equal(t, s.Id, cookie.Value)
}

func TestSession_Expired(t *testing.T) {
	t.Parallel()
	u, _ := models.NewUser("Batman", "batman@hotmail.com", "password123", "")
	s := models.NewSession(u)

	assert.False(t, s.Expired())
}
