package env

import "os"

func NewEnv() *Env {
	return &Env{}
}

type Env struct{}

func (e *Env) Get(key string) string {
	env := os.Getenv(key)
	if env != "" {
		return env
	}
	return ""
}
