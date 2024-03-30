package models

import "errors"

var (
	ErrEmptyID  = errors.New("ID is empty")
	ErrNotFound = errors.New("item not found")
)
