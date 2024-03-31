package main

import (
	"log/slog"
	"os"

	"github.com/Linkinlog/loggr/internal/handlers"
	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/stores"
)

const addr = ":8080"

func main() {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	b := &models.Garden{
		Name:        "Bandit Acres",
		Location:    "Somewhere",
		Description: "Worlds sweetest boy",
		Image: &models.Image{
			URL:       "https://i.ibb.co/qR0Qz05/image.jpg",
			Thumbnail: "https://i.ibb.co/qR0Qz05/image.jpg",
			Id:        "qR0Qz05",
		},
	}
	ma := &models.Garden{
		Name:        "Maggie Falls",
		Location:    "All around",
		Description: "The bestest girl",
		Image: &models.Image{
			URL:       "https://i.ibb.co/HCGb3p4/IMG-5361.jpg",
			Thumbnail: "https://i.ibb.co/HCGb3p4/IMG-5361.jpg",
			Id:        "HCGb3p4",
		},
	}
	mi := &models.Garden{
		Name:        "Mimarigolds",
		Location:    "Blooms everywhere",
		Description: "Needs direct sunlight And plenty of attention",
		Image: &models.Image{
			URL:       "https://i.ibb.co/G7dpvmZ/IMG-2115.jpg",
			Thumbnail: "https://i.ibb.co/G7dpvmZ/IMG-2115.jpg",
			Id:        "G7dpvmZ",
		},
	}

	m := stores.NewInMemory([]*models.Garden{b, ma, mi})
	ssr := handlers.NewSSR(l, addr, m)

	l.Info("its_alive!", slog.String("addr", addr))

	if err := ssr.ServeHTTP(); err != nil {
		l.Error("error serving http", err)
	}
}
