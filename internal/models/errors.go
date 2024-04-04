package models

import "errors"

var (
	ErrEmptyID              = errors.New("ID is empty")
	ErrAlreadyExists        = errors.New("item already exists")
	ErrNotFound             = errors.New("item not found")
	PwTooShort              = errors.New("password must be at least 8 characters long")
	NoUserInContext         = errors.New("no user in context")
	GardenAlreadyRegistered = errors.New("garden already registered")
)
