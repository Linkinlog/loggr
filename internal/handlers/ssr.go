package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/Linkinlog/loggr/assets"
	"github.com/Linkinlog/loggr/internal/env"
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
)

type wrapper struct {
	http.ResponseWriter
	s int
}

func (w *wrapper) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.s = statusCode
}

func NewSSR(l *slog.Logger, a string, s repositories.GardenStorer, u repositories.UserStorer) *SSR {
	return &SSR{
		logger:   l,
		addr:     a,
		g:        s,
		u:        u,
		sessions: make(map[string]*models.Session),
	}
}

type SSR struct {
	logger   *slog.Logger
	addr     string
	g        repositories.GardenStorer
	u        repositories.UserStorer
	sessions map[string]*models.Session
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
				s.logger.Error("error getting user from request form", "error", uErr.Error())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}
		ctx := u.ToContext(r.Context())
		err := handler(wr, r.WithContext(ctx))
		execTime := time.Since(start)
		if err != nil {
			s.logger.Error("error handling request", "error", err.Error(), "time", execTime)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		s.logger.Info("hit", "status", wr.s, "method", r.Method, "path", r.URL.Path, "time", execTime.String())
	}
}

func (s *SSR) serveGardens() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.wrapHandler(s.handleGardenListing))
	mux.HandleFunc("POST /", s.wrapHandler(s.handleNewGarden))
	mux.HandleFunc("GET /new", s.wrapHandler(handleNewGardenForm))

	mux.HandleFunc("GET /{id}", s.wrapHandler(s.handleGarden))
	mux.HandleFunc("POST /{id}", s.wrapHandler(s.handleUpdateGarden))
	mux.HandleFunc("GET /{id}/edit", s.wrapHandler(s.handleEditGardenForm))
	mux.HandleFunc("GET /{id}/delete", s.wrapHandler(s.handleDeleteGarden))

	mux.HandleFunc("GET /{id}/inventory", s.wrapHandler(s.handleGardenInventory))
	mux.HandleFunc("POST /{id}/inventory", s.wrapHandler(s.handleNewGardenInventoryItem))
	mux.HandleFunc("GET /{id}/inventory/new", s.wrapHandler(s.handleNewGardenInventoryItemForm))
	mux.HandleFunc("GET /{id}/inventory/{itemId}/edit", s.wrapHandler(s.handleEditGardenInventoryItemForm))
	mux.HandleFunc("GET /{id}/inventory/{itemId}/delete", s.wrapHandler(s.handleDeleteGardenInventoryItem))
	mux.HandleFunc("GET /{id}/inventory/{itemId}", s.wrapHandler(s.handleGardenInventoryItem))
	mux.HandleFunc("POST /{id}/inventory/{itemId}", s.wrapHandler(s.handleUpdateGardenInventoryItem))

	return mux
}

func (s *SSR) serveAuth() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.wrapHandler(handleNotFound))
	mux.HandleFunc("GET /sign-in", s.wrapHandler(handleSignInPage))
	mux.HandleFunc("POST /sign-in", s.wrapHandler(s.handleSignIn))
	mux.HandleFunc("GET /sign-out", s.wrapHandler(s.handleSignOut))
	mux.HandleFunc("GET /sign-up", s.wrapHandler(handleSignUpPage))
	mux.HandleFunc("POST /sign-up", s.wrapHandler(s.handleSignUp))
	mux.HandleFunc("GET /forgot-password", s.wrapHandler(handleForgotPassword))
	mux.HandleFunc("POST /forgot-password", s.wrapHandler(handleForgotPassword))

	return mux
}

func (s *SSR) getGarden(r *http.Request) (*models.Garden, error) {
	id := r.PathValue("id")
	repo := repositories.NewGardenRepository(s.g)
	g, err := repo.Get(id)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (s *SSR) handleDeleteGarden(w http.ResponseWriter, r *http.Request) error {
	g, err := s.getGarden(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
	}

	repo := repositories.NewGardenRepository(s.g)
	err = repo.Delete(g.Id())
	if err != nil {
		return err
	}

	http.Redirect(w, r, "/gardens/", http.StatusSeeOther)
	return nil
}

func (s *SSR) handleUpdateGarden(w http.ResponseWriter, r *http.Request) error {
	g, err := s.getGarden(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
	}

	name := r.FormValue("name")
	location := r.FormValue("location")
	description := r.FormValue("description")

	if name == "" || location == "" {
		return ErrNameAndLocationReq
	}

	img := g.Image

	imageFile, handler, err := r.FormFile("image")
	if err == nil {
		bbKey := env.NewEnv().Get("IMG_BB_KEY")
		var sErr error
		img, sErr = services.NewImageBB(bbKey).StoreImage(imageFile, handler.Filename)
		if sErr != nil {
			return sErr
		}
	}

	g.Name = name
	g.Location = location
	g.Description = description
	g.Image = img

	repo := repositories.NewGardenRepository(s.g)

	err = repo.Update(g.Id(), g)
	if err != nil {
		return err
	}

	http.Redirect(w, r, "/gardens/"+g.Id(), http.StatusSeeOther)
	return nil
}

func (s *SSR) handleEditGardenForm(w http.ResponseWriter, r *http.Request) error {
	g, err := s.getGarden(r)
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

func (s *SSR) handleDeleteGardenInventoryItem(w http.ResponseWriter, r *http.Request) error {
	g, err := s.getGarden(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
	}

	itemID := r.PathValue("itemId")
	repo := repositories.NewGardenRepository(s.g)
	err = repo.RemoveItemFromGarden(g.Id(), itemID)
	if err != nil {
		return err
	}

	http.Redirect(w, r, "/gardens/"+g.Id(), http.StatusSeeOther)
	return nil
}

func (s *SSR) handleUpdateGardenInventoryItem(w http.ResponseWriter, r *http.Request) error {
	g, err := s.getGarden(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
	}

	itemID := r.PathValue("itemId")
	repo := repositories.NewGardenRepository(s.g)
	item, err := repo.GetItemFromGarden(g.Id(), itemID)
	if err != nil {
		return err
	}
	if item == nil {
		return handleNotFound(w, r)
	}

	name := r.FormValue("name")
	t, _ := strconv.Atoi(r.FormValue("type"))
	fields := [5]*models.Field{
		models.NewField("field-1", r.FormValue("field-1")),
		models.NewField("field-2", r.FormValue("field-2")),
		models.NewField("field-3", r.FormValue("field-3")),
		models.NewField("field-4", r.FormValue("field-4")),
		models.NewField("field-5", r.FormValue("field-5")),
	}

	img := item.Image

	imageFile, handler, err := r.FormFile("image")
	if err == nil {
		bbKey := env.NewEnv().Get("IMG_BB_KEY")
		var sErr error
		img, sErr = services.NewImageBB(bbKey).StoreImage(imageFile, handler.Filename)
		if sErr != nil {
			return sErr
		}
	}

	item.Name = name
	item.Type = models.ItemType(t)
	item.Fields = fields
	item.Image = img

	err = repo.UpdateItemInGarden(g.Id(), itemID, item)
	if err != nil {
		return err
	}

	http.Redirect(w, r, "/gardens/"+g.Id()+"/inventory/"+itemID, http.StatusSeeOther)
	return nil
}

func (s *SSR) handleEditGardenInventoryItemForm(w http.ResponseWriter, r *http.Request) error {
	g, err := s.getGarden(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
		return err
	}

	itemID := r.PathValue("itemId")
	repo := repositories.NewGardenRepository(s.g)
	item, err := repo.GetItemFromGarden(g.Id(), itemID)
	if err != nil {
		return err
	}
	if item == nil {
		return handleNotFound(w, r)
	}

	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("Edit Inventory Item", "Welcome to the edit inventory item page", u)

	return p.Layout(web.EditGardenInventoryItemForm(g.Id(), item)).Render(r.Context(), w)
}

func (s *SSR) handleGardenInventoryItem(w http.ResponseWriter, r *http.Request) error {
	g, err := s.getGarden(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
		return err
	}

	itemID := r.PathValue("itemId")
	repo := repositories.NewGardenRepository(s.g)
	item, err := repo.GetItemFromGarden(g.Id(), itemID)
	if err != nil {
		return err
	}
	if item == nil {
		return handleNotFound(w, r)
	}

	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage(item.Name, "Welcome to the garden inventory item page", u)

	return p.Layout(web.GardenInventoryItem(g.Id(), item)).Render(r.Context(), w)
}

func (s *SSR) handleNewGardenInventoryItem(w http.ResponseWriter, r *http.Request) error {
	g, err := s.getGarden(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
		return err
	}

	name := r.FormValue("name")
	t, _ := strconv.Atoi(r.FormValue("type"))
	fields := [5]*models.Field{
		models.NewField("field-1", r.FormValue("field-1")),
		models.NewField("field-2", r.FormValue("field-2")),
		models.NewField("field-3", r.FormValue("field-3")),
		models.NewField("field-4", r.FormValue("field-4")),
		models.NewField("field-5", r.FormValue("field-5")),
	}

	img := models.NewImage("not-found", "/assets/imageNotFound.webp", "/assets/imageNotFound.webp", "")

	imageFile, handler, err := r.FormFile("image")
	if err == nil {
		bbKey := env.NewEnv().Get("IMG_BB_KEY")
		var sErr error
		img, sErr = services.NewImageBB(bbKey).StoreImage(imageFile, handler.Filename)
		if sErr != nil {
			return sErr
		}
	}

	i := models.NewItem(name, img, models.ItemType(t), fields)

	repo := repositories.NewGardenRepository(s.g)

	err = repo.AddItemToGarden(g.Id(), i)
	if err != nil {
		return err
	}

	http.Redirect(w, r, "/gardens/"+g.Id(), http.StatusSeeOther)
	return nil
}

func (s *SSR) handleNewGardenInventoryItemForm(w http.ResponseWriter, r *http.Request) error {
	g, err := s.getGarden(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
		return err
	}

	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("New Inventory Item", "Welcome to the new inventory item page", u)

	return p.Layout(web.NewGardenInventoryItemForm(g.Id())).Render(r.Context(), w)
}

func (s *SSR) handleGardenInventory(w http.ResponseWriter, r *http.Request) error {
	g, err := s.getGarden(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
		return err
	}

	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage(g.Name, "Welcome to the garden inventory page", u)

	return p.Layout(web.GardenInventory(g)).Render(r.Context(), w)
}

func (s *SSR) handleGarden(w http.ResponseWriter, r *http.Request) error {
	g, err := s.getGarden(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
		return err
	}

	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage(g.Name, "Welcome to the garden page", u)

	return p.Layout(web.Garden(g)).Render(r.Context(), w)
}

func (s *SSR) handleNewGarden(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return err
	}
	name := r.FormValue("name")
	location := r.FormValue("location")
	description := r.FormValue("description")

	img := models.NewImage("not-found", "/assets/imageNotFound.webp", "/assets/imageNotFound.webp", "")

	imageFile, handler, err := r.FormFile("image")
	if err == nil {
		bbKey := env.NewEnv().Get("IMG_BB_KEY")
		var sErr error
		img, sErr = services.NewImageBB(bbKey).StoreImage(imageFile, handler.Filename)
		if sErr != nil {
			return sErr
		}
	}

	g := models.NewGarden(name, location, description, img, []*models.Item{})

	repo := repositories.NewGardenRepository(s.g)

	_, err = repo.Add(g)
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
	u, _ := models.UserFromContext(r.Context())
	p := web.NewPage("Gardens", "Welcome to the gardens page", u)

	repo := repositories.NewGardenRepository(s.g)
	gardens, err := repo.List()
	if err != nil {
		return err
	}

	return p.Layout(web.GardenListing(gardens)).Render(r.Context(), w)
}

func (s *SSR) handleSignUp(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	repo := repositories.NewUserRepository(s.u)

	if _, err := repo.Get(email); err == nil {
		p := web.NewPage("Sign Up", "Welcome to the sign up page", nil)

		_ = p.Layout(web.SignUp(ErrUserExists.Error())).Render(r.Context(), w)
		return nil
	}

	u, err := models.NewUser(name, email, password)
	if err != nil {
		p := web.NewPage("Sign Up", "Welcome to the sign up page", nil)

		_ = p.Layout(web.SignUp(err.Error())).Render(r.Context(), w)
		return nil
	}

	_, err = repo.Add(u)
	if err != nil {
		p := web.NewPage("Sign Up", "Welcome to the sign up page", nil)

		_ = p.Layout(web.SignUp(err.Error())).Render(r.Context(), w)
		return nil
	}

	sess := models.NewSession(u)
	s.sessions[sess.Id()] = sess

	http.SetCookie(w, sess.ToCookie())

	http.Redirect(w, r, "/gardens/", http.StatusSeeOther)
	return nil
}

func (s *SSR) handleSignIn(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
		p := web.NewPage("Sign In", "Welcome to the sign in page", nil)

		_ = p.Layout(web.SignIn(ErrEmailAndPassReq.Error())).Render(r.Context(), w)
		return nil
	}

	repo := repositories.NewUserRepository(s.u)
	u, err := repo.Get(email)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			p := web.NewPage("Sign In", "Welcome to the sign in page", nil)

			_ = p.Layout(web.SignIn(models.ErrNotFound.Error())).Render(r.Context(), w)
			return nil
		}
		return err
	}

	if !u.CheckPassword(password) {
		p := web.NewPage("Sign In", "Welcome to the sign in page", nil)

		_ = p.Layout(web.SignIn(ErrorInvalidPassword.Error())).Render(r.Context(), w)
		return nil
	}
	sess := models.NewSession(u)
	s.sessions[sess.Id()] = sess

	http.SetCookie(w, sess.ToCookie())

	http.Redirect(w, r, "/gardens/", http.StatusSeeOther)
	return nil
}

func (s *SSR) handleSignOut(w http.ResponseWriter, r *http.Request) error {
	token, err := r.Cookie("token")
	if err != nil {
		return err
	}

	delete(s.sessions, token.Value)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
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
