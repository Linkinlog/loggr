package handlers

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/Linkinlog/loggr/assets"
	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/repositories"
	"github.com/Linkinlog/loggr/internal/services"
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
	ErrPassReq            = errors.New("password required")
	ErrInvalidCode        = errors.New("invalid code")
)

type wrapper struct {
	http.ResponseWriter
	s int
}

func (w *wrapper) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.s = statusCode
}

func NewSSR(
	l *slog.Logger,
	a string,
	d []*models.Garden,
	u *repositories.UserRepository,
	g *repositories.GardenRepository,
	i *repositories.ItemRepository,
	s *repositories.SessionRepository,
	ms *services.MailService,
) *SSR {
	return &SSR{
		logger:         l,
		addr:           a,
		defaultGardens: d,
		i:              i,
		u:              u,
		g:              g,
		s:              s,
		ms:             ms,
	}
}

type SSR struct {
	logger         *slog.Logger
	addr           string
	defaultGardens []*models.Garden
	u              *repositories.UserRepository
	g              *repositories.GardenRepository
	i              *repositories.ItemRepository
	s              *repositories.SessionRepository
	ms             *services.MailService
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

	sess, err := s.s.Get(token.Value)
	if err != nil {
		return nil, err
	}

	if sess.Expired() {
		err = s.s.Delete(sess.Id)
		if err != nil {
			return nil, err
		}
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
			if !errors.Is(uErr, http.ErrNoCookie) && uErr.Error() != "session not found" && !errors.Is(uErr, sql.ErrNoRows) {
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

	return p.Layout(web.EditGarden(g, "")).Render(r.Context(), w)
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
	item, err := s.i.Get(itemID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
		return err
	}
	if item == nil {
		return handleNotFound(w, r)
	}

	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("Edit Inventory Item", "Welcome to the edit inventory item page", u)

	return p.Layout(web.EditGardenInventoryItemForm(g.Id, item, "")).Render(r.Context(), w)
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

	items, err := s.g.GetItemsForGarden(g.Id)
	if err != nil {
		if !errors.Is(err, models.ErrNotFound) {
			return err
		}
	}
	for _, i := range items {
		if i.Id == itemID {
			p := web.NewPage(i.Name, "Welcome to the garden inventory item page", u)

			return p.Layout(web.GardenInventoryItem(g.Id, i)).Render(r.Context(), w)
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
		return err
	}

	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("New Inventory Item", "Welcome to the new inventory item page", u)

	return p.Layout(web.NewGardenInventoryItemForm(g.Id, "", "", "", "", "", "", 0, "")).Render(r.Context(), w)
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

	inventory, err := s.g.GetItemsForGarden(g.Id)
	if err != nil {
		if !errors.Is(err, models.ErrNotFound) {
			return err
		}
	}

	search := r.URL.Query().Has("search")
	query := ""
	if search {
		query = r.URL.Query().Get("search")
		http.Redirect(w, r, "/gardens/"+g.Id+"?search="+query, http.StatusSeeOther)
	}

	return p.Layout(web.GardenInventory(g.Id, query, inventory)).Render(r.Context(), w)
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

	items, err := s.g.GetItemsForGarden(g.Id)
	if err != nil {
		if !errors.Is(err, models.ErrNotFound) {
			return err
		}
	}

	search := r.URL.Query().Has("search")
	query := ""
	if search {
		query = r.URL.Query().Get("search")
		items = models.SearchItems(items, query)
	}

	return p.Layout(web.Garden(g, items, query)).Render(r.Context(), w)
}

func (s *SSR) handleGardenListing(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return handleNotFound(w, r)
	}
	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("Gardens", "Welcome to the gardens page", u)

	gardens := s.defaultGardens
	if u != nil {
		var err error
		gardens, err = s.u.GetGardensForUser(u.Id)
		if err != nil {
			if !errors.Is(err, models.ErrNotFound) {
				return err
			}
		}
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

	return p.Layout(web.NewGarden("", "", "", "")).Render(r.Context(), w)
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

func handleForgotPasswordForm(w http.ResponseWriter, r *http.Request) error {
	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("Forgot Password", "Welcome to the forgot password page", u)

	return p.Layout(web.ForgotPassword("", "")).Render(r.Context(), w)
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

func handleResetPasswordForm(w http.ResponseWriter, r *http.Request) error {
	code := r.PathValue("resetCode")
	p := web.NewPage("Reset Password", "Welcome to the reset password page", nil)

	return p.Layout(web.ResetPassword(code, "", "")).Render(r.Context(), w)
}
