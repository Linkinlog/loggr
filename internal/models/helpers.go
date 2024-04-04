package models

import (
	"crypto/rand"
	"fmt"
	"time"
)

const (
	MaxTTL         = 7 * 24 * time.Hour
	MinTokenLength = 32
)

func genId() string {
	n := 8
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", b)
}

func genToken(l int) string {
	if l < MinTokenLength {
		l = MinTokenLength
	}
	b := make([]byte, l)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", b)
}
