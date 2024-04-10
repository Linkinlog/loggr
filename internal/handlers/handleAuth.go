package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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
	mux.HandleFunc("GET /forgot-password", s.wrapHandler(handleForgotPasswordForm))
	mux.HandleFunc("POST /forgot-password", s.wrapHandler(s.handleForgotPassword))
	mux.HandleFunc("POST /reset-password/{resetCode}", s.wrapHandler(s.handleResetPassword))
	mux.HandleFunc("GET /reset-password/{resetCode}", s.wrapHandler(handleResetPasswordForm))

	return mux
}

func (s *SSR) handleResetPassword(w http.ResponseWriter, r *http.Request) error {
	code := r.PathValue("resetCode")
	password := r.FormValue("password")
	if password == "" || code == "" {
		p := web.NewPage("Reset Password", "Welcome to the reset password page", nil)

		_ = p.Layout(web.ResetPassword(code, "password required", "")).Render(r.Context(), w)
		return nil
	}

	u, err := s.u.GetByResetCode(code)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) || errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("invalid reset code")
		}
		p := web.NewPage("Reset Password", "Welcome to the reset password page", nil)

		_ = p.Layout(web.ResetPassword(code, err.Error(), "")).Render(r.Context(), w)
		return nil
	}

	if u != nil {
		if err := u.ChangePassword(password); err != nil {
			return err
		}
		if err := s.u.Update(u.Id, u); err != nil {
			return err
		}
	}

	p := web.NewPage("Reset Password", "Welcome to the reset password page", nil)

	_ = p.Layout(web.ResetPassword(code, "", "Password reset, please try it out")).Render(r.Context(), w)
	return s.u.ClearResetCode(u.Email)
}

func (s *SSR) handleForgotPassword(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	if email == "" {
		p := web.NewPage("Forgot Password", "Welcome to the forgot password page", nil)

		_ = p.Layout(web.ForgotPassword("email required", "")).Render(r.Context(), w)
		return nil
	}

	if _, err := s.u.GetByEmail(email); err != nil {
		p := web.NewPage("Forgot Password", "Welcome to the forgot password page", nil)

		_ = p.Layout(web.ForgotPassword("email not found", "")).Render(r.Context(), w)
		return nil
	}

	code, err := s.u.GenerateResetCode(email)
	if err != nil {
		return err
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	resetLink := url.URL{
		Scheme: scheme,
		Host:   r.Host,
		Path:   "/auth/reset-password/" + code,
	}

	if resp, err := s.ms.SendResetPassword(email, resetLink.String()); err != nil {
		err = fmt.Errorf("error sending email: %w, response %s", err, resp)
		return err
	}

	p := web.NewPage("Forgot Password", "Welcome to the forgot password page", nil)

	_ = p.Layout(web.ForgotPassword("", "Email sent! This code will expire shortly")).Render(r.Context(), w)

	time.AfterFunc(10*time.Minute, func() {
		_ = s.u.ClearResetCode(email)
	})
	return nil
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

	u, err := models.NewUser(name, email, password, "/assets/imageNotFound.webp")
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
