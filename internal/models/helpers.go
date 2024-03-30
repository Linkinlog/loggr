package models

import (
	"crypto/rand"
	"fmt"
)

func genId() string {
	n := 8
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", b)
}
