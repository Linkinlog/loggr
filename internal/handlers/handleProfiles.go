package handlers

import (
	"errors"
	"net/http"

	"github.com/Linkinlog/loggr/internal/env"
	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/services"
	"github.com/Linkinlog/loggr/web"
)

func (s *SSR) serveProfiles() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.wrapHandler(s.handleProfile))
	mux.HandleFunc("POST /", s.wrapHandler(s.handleUpdateProfile))
	mux.HandleFunc("GET /edit", s.wrapHandler(s.handleEditProfileForm))
	mux.HandleFunc("GET /delete", s.wrapHandler(s.handleDeleteProfile))
	return mux
}

func (s *SSR) handleProfile(w http.ResponseWriter, r *http.Request) error {
	u, err := models.UserFromContext(r.Context())
	if err != nil {
		if errors.Is(err, models.NoUserInContext) {
			return handleLanding(w, r)
		}
		return err
	}
	p := web.NewPage("Profile Page", "Welcome to the profile page", u)

	return p.Layout(web.Profile(*u)).Render(r.Context(), w)
}

func (s *SSR) handleUpdateProfile(w http.ResponseWriter, r *http.Request) error {
	u, err := models.UserFromContext(r.Context())
	if err != nil {
		if errors.Is(err, models.NoUserInContext) {
			return handleLanding(w, r)
		}
		return err
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if name == "" || email == "" {
		return errors.New("name and email are required")
	}

	if foundUser, _ := s.u.GetByEmail(email); foundUser != nil && foundUser.Id != u.Id {
		p := web.NewPage("Edit Profile Page", "Welcome to the edit profile page", u)

		err := "email already in use"
		return p.Layout(web.EditProfileForm(*u, err)).Render(r.Context(), w)
	}

	if password == "" {
		password = string(u.Password)
	}

	img := u.Image

	imageFile, handler, err := r.FormFile("image")
	if err == nil {
		bbKey := env.NewEnv().GetOrDefault("IMG_BB_KEY", "")
		var sErr error
		img, sErr = services.NewImageBB(bbKey).StoreImage(imageFile, handler.Filename)
		if sErr != nil {
			if errors.Is(sErr, services.ErrImageUpload) {
				p := web.NewPage("Edit Profile Page", "Welcome to the edit profile page", u)

				err := "error uploading image, please try a different image"
				return p.Layout(web.EditProfileForm(*u, err)).Render(r.Context(), w)
			}
			return sErr
		}
	}

	newUser, err := models.NewUser(name, email, password, img)
	if err != nil {
		return err
	}

	if err := s.u.Update(u.Id, newUser); err != nil {
		return err
	}

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
	return nil
}

func (s *SSR) handleEditProfileForm(w http.ResponseWriter, r *http.Request) error {
	u, err := models.UserFromContext(r.Context())
	if err != nil {
		if errors.Is(err, models.NoUserInContext) {
			return handleLanding(w, r)
		}
		return err
	}
	p := web.NewPage("Edit Profile Page", "Welcome to the edit profile page", u)

	return p.Layout(web.EditProfileForm(*u, "")).Render(r.Context(), w)
}

func (s *SSR) handleDeleteProfile(w http.ResponseWriter, r *http.Request) error {
	u, err := models.UserFromContext(r.Context())
	if err != nil {
		if errors.Is(err, models.NoUserInContext) {
			return handleLanding(w, r)
		}
		return err
	}

	if err := s.u.Delete(u.Id); err != nil {
		return err
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}
