package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/web"
)

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

func (s *SSR) handleSignOut(w http.ResponseWriter, r *http.Request) error {
	token, err := r.Cookie("token")
	if err != nil {
		return err
	}

	err = s.s.Delete(token.Value)
	if err != nil {
		return err
	}

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

func (s *SSR) handleSignIn(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
		p := web.NewPage("Sign In", "Welcome to the sign in page", nil)

		_ = p.Layout(web.SignIn(ErrEmailAndPassReq.Error())).Render(r.Context(), w)
		return nil
	}

	u, err := s.u.GetByEmail(email)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) || errors.Is(err, sql.ErrNoRows) {
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
	err = s.s.Add(sess)
	if err != nil {
		return err
	}

	http.SetCookie(w, sess.ToCookie())

	http.Redirect(w, r, "/gardens/", http.StatusSeeOther)
	return nil
}

func (s *SSR) handleSignUp(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if _, err := s.u.GetByEmail(email); err == nil {
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

	_, err = s.u.Add(u)
	if err != nil {
		p := web.NewPage("Sign Up", "Welcome to the sign up page", nil)

		_ = p.Layout(web.SignUp(err.Error())).Render(r.Context(), w)
		return nil
	}

	sess := models.NewSession(u)
	err = s.s.Add(sess)
	if err != nil {
		return err
	}

	http.SetCookie(w, sess.ToCookie())

	http.Redirect(w, r, "/gardens/", http.StatusSeeOther)
	return nil
}
