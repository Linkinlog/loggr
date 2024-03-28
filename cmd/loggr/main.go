package main

import (
	"log/slog"
	"os"

	"github.com/Linkinlog/loggr/internal/handlers"
)

const addr = ":1420"

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ssr := handlers.NewSSR(l, addr)

	l.Info("starting server", slog.String("addr", addr))

	if err := ssr.ServeHTTP(); err != nil {
		l.Error("error serving http", err)
	}
}
