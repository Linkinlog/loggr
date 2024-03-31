package models

func NewImage(id, url string) *Image {
	return &Image{
		Id:  id,
		URL: url,
	}
}

type Image struct {
	Id        string
	URL       string
	Thumbnail string
	DeleteURL string
}
