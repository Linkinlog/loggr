package models

import "errors"

var (
	ErrEmptyID      = errors.New("ID is empty")
	ErrNotFound     = errors.New("item not found")
	PwTooShort      = errors.New("password must be at least 8 characters long")
	NoUserInContext = errors.New("no user in context")
)
