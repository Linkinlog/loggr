package models

import (
	"net/http"
	"time"
)

func NewSession(u *User) *Session {
	return &Session{
		id:        genToken(MinTokenLength),
		User:      u,
		expiresAt: time.Now().Add(MaxTTL),
	}
}

type Session struct {
	id        string
	User      *User
	expiresAt time.Time
}

func (s *Session) Id() string {
	return s.id
}

func (s *Session) TTL() time.Duration {
	return time.Until(s.expiresAt)
}

func (s *Session) ToCookie() *http.Cookie {
	return &http.Cookie{
		Name:     "token",
		Value:    s.Id(),
		MaxAge:   int(s.TTL().Seconds()),
		Expires:  s.expiresAt,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}
}

func (s *Session) Expired() bool {
	return time.Since(s.expiresAt) > 0
}
