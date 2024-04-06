package models

import (
	"net/http"
	"time"
)

func NewSession(u *User) *Session {
	if u == nil {
		return nil
	}
	return &Session{
		Id:        genToken(MinTokenLength),
		User:      u,
		expiresAt: time.Now().Add(MaxTTL),
	}
}

type Session struct {
	Id        string `db:"id"`
	User      *User
	expiresAt time.Time
}

func (s *Session) TTL() time.Duration {
	return time.Until(s.expiresAt)
}

func (s *Session) ToCookie() *http.Cookie {
	return &http.Cookie{
		Name:     "token",
		Value:    s.Id,
		MaxAge:   int(s.TTL().Seconds()),
		Expires:  s.expiresAt,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}
}

func (s *Session) Expired() bool {
	return time.Since(s.expiresAt) > 0
}
