package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/Linkinlog/loggr/assets"
	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/repositories"
	"github.com/Linkinlog/loggr/web"
	"github.com/a-h/templ"
)

var (
	ErrSessionNotFound    = errors.New("session not found")
	ErrSessionExpired     = errors.New("session expired")
	ErrorInternalServer   = errors.New("Internal Server Error")
	ErrNameAndLocationReq = errors.New("name and location are required")
	ErrEmailAndPassReq    = errors.New("email and password are required")
	ErrUserExists         = errors.New("user already exists")
	ErrorInvalidPassword  = errors.New("invalid password")
)

type wrapper struct {
	http.ResponseWriter
	s int
}

func (w *wrapper) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.s = statusCode
}

func NewSSR(l *slog.Logger, a string, d []*models.Garden, u repositories.UserStorer) *SSR {
	return &SSR{
		logger:         l,
		addr:           a,
		defaultGardens: d,
		u:              u,
		sessions:       make(map[string]*models.Session),
	}
}

type SSR struct {
	logger         *slog.Logger
	addr           string
	defaultGardens []*models.Garden
	u              repositories.UserStorer
	sessions       map[string]*models.Session
}

func (s *SSR) ServeHTTP() error {
	mux := http.NewServeMux()

	mux.Handle("GET /", s.wrapHandler(handleLanding))
	mux.Handle("GET /auth/", http.StripPrefix("/auth", s.serveAuth()))
	mux.Handle("POST /auth/", http.StripPrefix("/auth", s.serveAuth()))
	mux.Handle("GET /gardens/", http.StripPrefix("/gardens", s.serveGardens()))
	mux.Handle("POST /gardens/", http.StripPrefix("/gardens", s.serveGardens()))
	mux.Handle("GET /about", s.wrapHandler(handleAbout))

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assets.NewAssets()))))

	// https://templ.guide/syntax-and-usage/css-style-management/#css-middleware
	handler := templ.NewCSSMiddleware(mux, web.Styles()...)

	server := &http.Server{
		Addr:    s.addr,
		Handler: handler,
	}

	return server.ListenAndServe()
}

func (s *SSR) userFromRequest(r *http.Request) (*models.User, error) {
	token, err := r.Cookie("token")
	if err != nil {
		return nil, err
	}

	sess, ok := s.sessions[token.Value]
	if !ok {
		return nil, ErrSessionNotFound
	}

	if sess.Expired() {
		delete(s.sessions, token.Value)
		return nil, ErrSessionExpired
	}

	u := sess.User

	return u, nil
}

func (s *SSR) wrapHandler(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wr := &wrapper{ResponseWriter: w}
		u, uErr := s.userFromRequest(r)
		if uErr != nil {
			if !errors.Is(uErr, http.ErrNoCookie) && uErr.Error() != "session not found" {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				s.logger.Error("error getting user from request form", "error", uErr.Error())
				return
			}
		}
		ctx := u.ToContext(r.Context())
		err := handler(wr, r.WithContext(ctx))
		execTime := time.Since(start)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			s.logger.Error("error handling request", "error", err.Error(), "time", execTime)
		}
		s.logger.Info("hit", "status", wr.s, "method", r.Method, "path", r.URL.Path, "time", execTime.String())
	}
}

func (s *SSR) handleEditGardenForm(w http.ResponseWriter, r *http.Request) error {
	g, _, err := s.getGardenForUser(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
		return err
	}

	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("Edit Garden", "Welcome to the edit garden page", u)

	return p.Layout(web.EditGarden(g)).Render(r.Context(), w)
}

func (s *SSR) handleEditGardenInventoryItemForm(w http.ResponseWriter, r *http.Request) error {
	g, _, err := s.getGardenForUser(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
		return err
	}

	itemID := r.PathValue("itemId")
	item := g.GetItem(itemID)
	if item == nil {
		return handleNotFound(w, r)
	}

	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("Edit Inventory Item", "Welcome to the edit inventory item page", u)

	return p.Layout(web.EditGardenInventoryItemForm(g.Id(), item)).Render(r.Context(), w)
}

func (s *SSR) handleGardenInventoryItem(w http.ResponseWriter, r *http.Request) error {
	g, u, err := s.getGardenForUser(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
		return err
	}

	itemID := r.PathValue("itemId")

	for _, i := range g.Inventory {
		if i.Id() == itemID {
			p := web.NewPage(i.Name, "Welcome to the garden inventory item page", u)

			return p.Layout(web.GardenInventoryItem(g.Id(), i)).Render(r.Context(), w)
		}
	}

	return handleNotFound(w, r)
}

func (s *SSR) handleNewGardenInventoryItemForm(w http.ResponseWriter, r *http.Request) error {
	g, _, err := s.getGardenForUser(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
		fmt.Println(err)
		return err
	}

	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("New Inventory Item", "Welcome to the new inventory item page", u)

	return p.Layout(web.NewGardenInventoryItemForm(g.Id())).Render(r.Context(), w)
}

func (s *SSR) handleGardenInventory(w http.ResponseWriter, r *http.Request) error {
	g, _, err := s.getGardenForUser(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
		return err
	}

	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage(g.Name, "Welcome to the garden inventory page", u)

	inventory := g.Inventory
	search := r.URL.Query().Has("search")
	query := ""
	if search {
		query = r.URL.Query().Get("search")
		http.Redirect(w, r, "/gardens/"+g.Id()+"?search="+query, http.StatusSeeOther)
	}

	return p.Layout(web.GardenInventory(g.Id(), query, inventory)).Render(r.Context(), w)
}

func (s *SSR) handleGarden(w http.ResponseWriter, r *http.Request) error {
	g, _, err := s.getGardenForUser(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
		return err
	}

	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage(g.Name, "Welcome to the garden page", u)

	gardens := g.Inventory

	search := r.URL.Query().Has("search")
	query := ""
	if search {
		query = r.URL.Query().Get("search")
		gardens = models.SearchItems(gardens, query)
	}

	return p.Layout(web.Garden(g, gardens, query)).Render(r.Context(), w)
}

func (s *SSR) handleGardenListing(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return handleNotFound(w, r)
	}
	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("Gardens", "Welcome to the gardens page", u)

	gardens := s.defaultGardens
	if u != nil {
		gardens = u.ListGardens()
	}

	search := r.URL.Query().Has("search")
	query := ""
	if search {
		query = r.URL.Query().Get("search")
		gardens = models.SearchGardens(gardens, query)
	}

	return p.Layout(web.GardenListing(gardens, query)).Render(r.Context(), w)
}

func handleNewGardenForm(w http.ResponseWriter, r *http.Request) error {
	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("New Garden", "Welcome to the new garden page", u)

	return p.Layout(web.NewGarden()).Render(r.Context(), w)
}

func handleSignInPage(w http.ResponseWriter, r *http.Request) error {
	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("Sign In", "Welcome to the sign in page", u)

	return p.Layout(web.SignIn("")).Render(r.Context(), w)
}

func handleSignUpPage(w http.ResponseWriter, r *http.Request) error {
	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("Sign Up", "Welcome to the sign up page", u)

	return p.Layout(web.SignUp("")).Render(r.Context(), w)
}

func handleForgotPassword(w http.ResponseWriter, r *http.Request) error {
	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("Forgot Password", "Welcome to the forgot password page", u)

	return p.Layout(web.ForgotPassword()).Render(r.Context(), w)
}

func handleLanding(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return handleNotFound(w, r)
	}
	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("Landing", "Welcome to the landing page", u)

	return p.Layout(web.Landing()).Render(r.Context(), w)
}

func handleAbout(w http.ResponseWriter, r *http.Request) error {
	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("About", "Welcome to the about page", u)

	return p.Layout(web.About()).Render(r.Context(), w)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) error {
	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("404", "Page not found", u)

	w.WriteHeader(http.StatusNotFound)
	return p.Layout(web.NotFoundPage()).Render(r.Context(), w)
}
