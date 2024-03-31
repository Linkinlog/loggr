package models

func NewImage(id, url, thumb, deleteUrl string) *Image {
	return &Image{
		Id:        id,
		URL:       url,
		Thumbnail: thumb,
		DeleteURL: deleteUrl,
	}
}

type Image struct {
	Id        string
	URL       string
	Thumbnail string
	DeleteURL string
}
