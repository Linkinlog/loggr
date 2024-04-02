package models

import (
	"time"
)

const (
	MaxTTL         = 7 * 24 * time.Hour
	MinTokenLength = 32
)

func NewToken(l int, ttl time.Duration) *Token {
	tokenId := genToken(l)
	if ttl > MaxTTL {
		ttl = MaxTTL
	}
	return &Token{
		Value: tokenId,
		TTL:   ttl,
		set:   time.Now(),
	}
}

type Token struct {
	Value string
	TTL   time.Duration
	set   time.Time
}

func (t *Token) String() string {
	return t.Value
}

func (t *Token) Expired() bool {
	return time.Since(t.set) > t.TTL
}
