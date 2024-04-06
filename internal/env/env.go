package env

import "os"

func NewEnv() *Env {
	return &Env{}
}

type Env struct{}

func (e *Env) GetOrDefault(key, def string) string {
	env := os.Getenv(key)
	if env != "" {
		return env
	}
	return def
}
