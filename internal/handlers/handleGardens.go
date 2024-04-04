package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Linkinlog/loggr/internal/env"
	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/repositories"
	"github.com/Linkinlog/loggr/internal/services"
)

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
	mux.HandleFunc("GET /{id}/inventory/{itemId}", s.wrapHandler(s.handleGardenInventoryItem))
	mux.HandleFunc("POST /{id}/inventory/{itemId}", s.wrapHandler(s.handleUpdateGardenInventoryItem))
	mux.HandleFunc("GET /{id}/inventory/{itemId}/edit", s.wrapHandler(s.handleEditGardenInventoryItemForm))
	mux.HandleFunc("GET /{id}/inventory/{itemId}/delete", s.wrapHandler(s.handleDeleteGardenInventoryItem))

	return mux
}

func (s *SSR) getGardenForUser(r *http.Request) (*models.Garden, *models.User, error) {
	id := r.PathValue("id")
	if id == "" {
		return nil, nil, models.ErrNotFound
	}

	gardens := s.defaultGardens
	u, err := s.userFromRequest(r)
	if err == nil {
		gardens = u.ListGardens()
	}

	for _, g := range gardens {
		if g.Id() == id {
			return g, u, nil
		}
	}

	return nil, nil, models.ErrNotFound
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

	u, err := models.UserFromContext(r.Context())
	if err == nil && u != nil {
		rErr := u.RegisterGarden(g)
		if rErr != nil {
			return rErr
		}
	} else {
		s.defaultGardens = append(s.defaultGardens, g)
	}

	http.Redirect(w, r, "/gardens/", http.StatusSeeOther)
	return nil
}

func (s *SSR) handleUpdateGarden(w http.ResponseWriter, r *http.Request) error {
	g, u, err := s.getGardenForUser(r)
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

	if u != nil {
		u.Gardens[g.Id()] = g

		repo := repositories.NewUserRepository(s.u)
		_, err := repo.Save(u)
		if err != nil {
			s.logger.Error("error saving user", "error", err.Error())
		}
	} else {
		for i, garden := range s.defaultGardens {
			if garden.Id() == g.Id() {
				s.defaultGardens[i] = g
			}
		}
	}

	http.Redirect(w, r, "/gardens/"+g.Id(), http.StatusSeeOther)
	return nil
}

func (s *SSR) handleDeleteGarden(w http.ResponseWriter, r *http.Request) error {
	g, u, err := s.getGardenForUser(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
	}

	if u == nil {
		for i, garden := range s.defaultGardens {
			if garden.Id() == g.Id() {
				s.defaultGardens = append(s.defaultGardens[:i], s.defaultGardens[i+1:]...)
			}
		}
	} else {
		_ = u.UnregisterGarden(g)
	}

	http.Redirect(w, r, "/gardens/", http.StatusSeeOther)
	return nil
}

func (s *SSR) handleNewGardenInventoryItem(w http.ResponseWriter, r *http.Request) error {
	g, u, err := s.getGardenForUser(r)
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
	g.AddItem(i)

	if u != nil {
		repo := repositories.NewUserRepository(s.u)
		_, err := repo.Save(u)
		if err != nil {
			s.logger.Error("error saving user", "error", err.Error())
		}
	}

	http.Redirect(w, r, "/gardens/"+g.Id(), http.StatusSeeOther)
	return nil
}

func (s *SSR) handleUpdateGardenInventoryItem(w http.ResponseWriter, r *http.Request) error {
	g, u, err := s.getGardenForUser(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
	}

	itemID := r.PathValue("itemId")
	item := g.GetItem(itemID)
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

	if u != nil {
		repo := repositories.NewUserRepository(s.u)
		_, err := repo.Save(u)
		if err != nil {
			s.logger.Error("error saving user", "error", err.Error())
		}
	}

	http.Redirect(w, r, "/gardens/"+g.Id()+"/inventory/"+itemID, http.StatusSeeOther)
	return nil
}

func (s *SSR) handleDeleteGardenInventoryItem(w http.ResponseWriter, r *http.Request) error {
	g, u, err := s.getGardenForUser(r)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return handleNotFound(w, r)
		}
	}

	itemID := r.PathValue("itemId")

	g.RemoveItem(g.GetItem(itemID))

	if u != nil {
		repo := repositories.NewUserRepository(s.u)
		_, err := repo.Save(u)
		if err != nil {
			s.logger.Error("error saving user", "error", err.Error())
		}
	}

	http.Redirect(w, r, "/gardens/"+g.Id(), http.StatusSeeOther)
	return nil
}
