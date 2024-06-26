package models_test

import (
	"context"
	"testing"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		name     string
		email    string
		password string
		wantErr  bool
	}{
		"valid": {
			name:     "Johnny Bravo",
			email:    "jbravo+coolDude@hotmail.com",
			password: "password123",
			wantErr:  false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := models.NewUser(tc.name, tc.email, tc.password, "")
			if (err != nil) != tc.wantErr {
				t.Fatalf("NewUser(%q, %q, %q) = %v, wantErr %v", tc.name, tc.email, tc.password, err, tc.wantErr)
			}
		})
	}
}

func TestUser_UserFromContext(t *testing.T) {
	t.Parallel()
	u, err := models.NewUser("Batman", "batman@hotmail.com", "password123", "")
	assert.Nil(t, err)

	user, err := models.UserFromContext(context.WithValue(context.Background(), models.UserCtxKey("user"), u))

	assert.Nil(t, err)
	assert.Equal(t, u, user)
}

func TestUser_ToContext(t *testing.T) {
	t.Parallel()
	u, err := models.NewUser("Batman", "batman@hotmail.com", "password123", "")
	assert.Nil(t, err)

	ctx := u.ToContext(context.Background())

	assert.NotNil(t, ctx)
	assert.Equal(t, u, ctx.Value(models.UserCtxKey("user")))
}

func TestUser_String(t *testing.T) {
	t.Parallel()
	u, err := models.NewUser("Batman Forever", "batman@hotmail.com", "password123", "")
	assert.Nil(t, err)

	assert.Equal(t, "Batman Forever", u.String())
}

func TestUser_Id(t *testing.T) {
	t.Parallel()
	u, err := models.NewUser("Batman Forever", "batman@hotmail.com", "password123", "")
	assert.Nil(t, err)

	assert.NotEmpty(t, u.Id)
}

func TestUser_Password(t *testing.T) {
	t.Parallel()
	u, err := models.NewUser("Batman Forever", "batman@hotmail.com", "password123", "")
	assert.Nil(t, err)

	assert.NotEmpty(t, u.Password)
}

func TestUser_CheckPassword(t *testing.T) {
	t.Parallel()
	password := "password123"
	u, err := models.NewUser("Batman", "batman@hotmail.com", password, "")
	if err != nil {
		t.Fatalf("NewUser failed: %v", err)
	}

	cases := map[string]struct {
		password string
		want     bool
	}{
		"correct password": {
			password: password,
			want:     true,
		},
		"incorrect password": {
			password: "wrongPassword",
			want:     false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := u.CheckPassword(tc.password)
			if got != tc.want {
				t.Fatalf("CheckPassword(%q) = %v, want %v", tc.password, got, tc.want)
			}
		})
	}
}

func TestUser_ChangePassword(t *testing.T) {
	t.Parallel()
	u, err := models.NewUser("Batman Forever", "batman@hotmail.com", "password123", "")
	assert.Nil(t, err)

	newPassword := "newPassword123"

	err = u.ChangePassword(newPassword)
	assert.Nil(t, err)
	assert.True(t, u.CheckPassword(newPassword))
}
