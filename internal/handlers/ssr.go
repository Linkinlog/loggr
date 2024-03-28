package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/Linkinlog/loggr/assets"
	"github.com/Linkinlog/loggr/web"
)

func NewSSR(l *slog.Logger, a string) *SSR {
	return &SSR{
		logger: l,
		addr:   a,
	}
}

type SSR struct {
	logger *slog.Logger
	addr   string
}

func (s *SSR) ServeHTTP() error {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assets.NewAssets()))))

	mux.HandleFunc("GET /about", s.wrapHandler(handleAbout))
	mux.HandleFunc("GET /", s.wrapHandler(handleLanding))

	server := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	return server.ListenAndServe()
}

func (s *SSR) wrapHandler(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			s.logger.Error("error handling request", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func handleLanding(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		// TODO we didnt add a 404 page :/
		http.NotFound(w, r)
		return nil
	}
	p := web.NewPage("Landing", "Welcome to the landing page")

	return p.Layout(web.Landing()).Render(context.Background(), w)
}

func handleAbout(w http.ResponseWriter, _ *http.Request) error {
	p := web.NewPage("About", "Welcome to the about page")

	return p.Layout(web.About()).Render(context.Background(), w)
}
