package handlers

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/Linkinlog/loggr/assets"
	"github.com/Linkinlog/loggr/internal/env"
	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/repositories"
	"github.com/Linkinlog/loggr/internal/services"
	"github.com/Linkinlog/loggr/web"
)

type wrapper struct {
	http.ResponseWriter
	s int
}

func (w *wrapper) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.s = statusCode
}

func NewSSR(l *slog.Logger, a string, s repositories.Storer) *SSR {
	return &SSR{
		logger: l,
		addr:   a,
		s:      s,
	}
}

type SSR struct {
	logger *slog.Logger
	addr   string
	s      repositories.Storer
}

func (s *SSR) wrapHandler(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wr := &wrapper{ResponseWriter: w}
		err := handler(wr, r)
		execTime := time.Since(start)
		if err != nil {
			s.logger.Error("error handling request", err, "time", execTime)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		s.logger.Info("hit", "status", wr.s, "method", r.Method, "path", r.URL.Path, "time", execTime.String())
	}
}

func (s *SSR) ServeHTTP() error {
	mux := http.NewServeMux()

	mux.Handle("GET /", s.wrapHandler(handleLanding))
	mux.Handle("GET /auth/", http.StripPrefix("/auth", s.serveAuth()))
	mux.Handle("GET /gardens/", http.StripPrefix("/gardens", s.serveGardens()))
	mux.Handle("POST /gardens/", http.StripPrefix("/gardens", s.serveGardens()))
	mux.Handle("GET /about", s.wrapHandler(handleAbout))

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assets.NewAssets()))))

	server := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	return server.ListenAndServe()
}

func (s *SSR) serveGardens() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.wrapHandler(s.handleGardenListing))
	mux.HandleFunc("POST /", s.wrapHandler(s.handleNewGarden))
	mux.HandleFunc("GET /new", s.wrapHandler(handleNewGardenForm))

	return mux
}

func (s *SSR) serveAuth() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.wrapHandler(handleNotFound))
	mux.HandleFunc("GET /sign-in", s.wrapHandler(handleSignIn))
	mux.HandleFunc("GET /sign-out", s.wrapHandler(handleSignOut))
	mux.HandleFunc("GET /sign-up", s.wrapHandler(handleSignUp))
	mux.HandleFunc("GET /forgot-password", s.wrapHandler(handleForgotPassword))

	return mux
}

func (s *SSR) handleNewGarden(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return err
	}
	name := r.FormValue("name")
	location := r.FormValue("location")
	description := r.FormValue("description")
	imageFile, handler, err := r.FormFile("image")
	if err != nil {
		return err
	}

	bbKey := env.NewEnv().Get("IMG_BB_KEY")
	img, err := services.NewImageBB(bbKey).StoreImage(imageFile, handler.Filename)
	if err != nil {
		return err
	}

	g := models.NewGarden(name, location, description, img, []*models.Item{})

	repo := repositories.NewGardenRepository(s.s)

	_, err = repo.AddGarden(g)
	if err != nil {
		return err
	}

	http.Redirect(w, r, "/gardens/", http.StatusSeeOther)
	return nil
}

func (s *SSR) handleGardenListing(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return handleNotFound(w, r)
	}
	p := web.NewPage("Gardens", "Welcome to the gardens page")

	repo := repositories.NewGardenRepository(s.s)
	gardens, err := repo.ListGardens()
	if err != nil {
		return err
	}

	return p.Layout(web.GardenListing(gardens)).Render(context.Background(), w)
}

func handleNewGardenForm(w http.ResponseWriter, _ *http.Request) error {
	p := web.NewPage("New Garden", "Welcome to the new garden page")

	return p.Layout(web.NewGarden()).Render(context.Background(), w)
}

func handleSignIn(w http.ResponseWriter, _ *http.Request) error {
	p := web.NewPage("Sign In", "Welcome to the sign in page")

	return p.Layout(web.SignIn()).Render(context.Background(), w)
}

func handleSignOut(w http.ResponseWriter, _ *http.Request) error {
	// TODO
	return errors.New("not implemented")
}

func handleSignUp(w http.ResponseWriter, _ *http.Request) error {
	p := web.NewPage("Sign Up", "Welcome to the sign up page")

	return p.Layout(web.SignUp()).Render(context.Background(), w)
}

func handleForgotPassword(w http.ResponseWriter, _ *http.Request) error {
	p := web.NewPage("Forgot Password", "Welcome to the forgot password page")

	return p.Layout(web.ForgotPassword()).Render(context.Background(), w)
}

func handleLanding(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return handleNotFound(w, r)
	}
	p := web.NewPage("Landing", "Welcome to the landing page")

	return p.Layout(web.Landing()).Render(context.Background(), w)
}

func handleAbout(w http.ResponseWriter, _ *http.Request) error {
	p := web.NewPage("About", "Welcome to the about page")

	return p.Layout(web.About()).Render(context.Background(), w)
}

func handleNotFound(w http.ResponseWriter, _ *http.Request) error {
	p := web.NewPage("404", "Page not found")

	w.WriteHeader(http.StatusNotFound)
	return p.Layout(web.NotFoundPage()).Render(context.Background(), w)
}
