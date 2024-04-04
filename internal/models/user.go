package models

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func NewUser(name, email string, passwordString string) (*User, error) {
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
		id:        genId(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		password:  password,
		Gardens:   make(map[string]*Garden),
	}, nil
}

func UserFromContext(ctx context.Context) (*User, error) {
	user, ok := ctx.Value(UserCtxKey("user")).(*User)
	if !ok {
		return nil, NoUserInContext
	}

	return user, nil
}

type User struct {
	id        string
	FirstName string
	LastName  string
	Email     string
	password  []byte
	Gardens   map[string]*Garden
}

type UserCtxKey string

func (u *User) ToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, UserCtxKey("user"), u)
}

func (u *User) String() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Password() []byte {
	return u.password
}

func (u *User) CheckPassword(passwordString string) bool {
	err := bcrypt.CompareHashAndPassword(u.password, []byte(passwordString))
	return err == nil
}

func (u *User) ChangePassword(passwordString string) error {
	password, err := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.password = password
	return nil
}

func (u *User) RegisterGarden(g *Garden) error {
	if _, ok := u.Gardens[g.Id()]; ok {
		return GardenAlreadyRegistered
	}

	u.Gardens[g.Id()] = g
	return nil
}

func (u *User) UnregisterGarden(g *Garden) error {
	if _, ok := u.Gardens[g.Id()]; !ok {
		return ErrNotFound
	}

	delete(u.Gardens, g.Id())
	return nil
}

func (u *User) ListGardens() []*Garden {
	gardens := make([]*Garden, 0, len(u.Gardens))
	for _, g := range u.Gardens {
		gardens = append(gardens, g)
	}
	return gardens
}
