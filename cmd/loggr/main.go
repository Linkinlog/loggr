package main

import (
	"log/slog"
	"os"

	"github.com/Linkinlog/loggr/internal/handlers"
)

const addr = ":8080"

func main() {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	ssr := handlers.NewSSR(l, addr)

	l.Info("its_alive!", slog.String("addr", addr))

	if err := ssr.ServeHTTP(); err != nil {
		l.Error("error serving http", err)
	}
}
