package models

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func NewUser(name, email, passwordString, image string) (*User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	var firstName, lastName string
	if len(strings.Split(name, " ")) > 1 {
		firstName = strings.Split(name, " ")[0]
		lastName = strings.Split(name, " ")[1]
	} else {
		firstName = name
		lastName = ""
	}

	return &User{
		Id:        genId(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Image:     image,
	}, nil
}

func UserFromContext(ctx context.Context) (*User, error) {
	user, ok := ctx.Value(UserCtxKey("user")).(*User)
	if !ok || user == nil {
		return nil, NoUserInContext
	}

	return user, nil
}

type User struct {
	Id        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Image     string `db:"image"`
	Password  []byte `db:"password"`
	ResetCode string `db:"reset_code"`
}

type UserCtxKey string

func (u *User) ToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, UserCtxKey("user"), u)
}

func (u *User) String() string {
	if u.LastName == "" {
		return u.FirstName
	}
	return u.FirstName + " " + u.LastName
}

func (u *User) CheckPassword(passwordString string) bool {
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(passwordString))
	return err == nil
}

func (u *User) ChangePassword(passwordString string) error {
	password, err := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.Password = password
	return nil
}
