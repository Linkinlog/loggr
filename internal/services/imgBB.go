package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/Linkinlog/loggr/internal/models"
)

var (
	ErrImageNil    = errors.New("imgbb: image is nil")
	ErrImageUpload = errors.New("imgbb: image upload failed")
)

func NewImageBB(k string) *ImageBB {
	u := &url.URL{
		Scheme:   "https",
		Host:     "api.imgbb.com",
		Path:     "/1/upload",
		RawQuery: "key=" + k,
	}
	return &ImageBB{
		apiUrl: u,
	}
}

type ImageBB struct {
	apiUrl *url.URL
}

func (i *ImageBB) StoreImage(image io.Reader, name string) (*models.Image, error) {
	if image == nil {
		return nil, ErrImageNil
	}
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	part, err := writer.CreateFormFile("image", name)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, image)
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		Method: "POST",
		URL:    i.apiUrl,
		Header: map[string][]string{
			"Content-Type": {writer.FormDataContentType()},
		},
		Body: io.NopCloser(buf),
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return parseResponse(resp)
}

func parseResponse(resp *http.Response) (*models.Image, error) {
	type imageBBResponse struct {
		Data struct {
			URL    string `json:"url"`
			Medium struct {
				URL string `json:"url"`
			} `json:"medium"`
			Thumb struct {
				URL string `json:"url"`
			} `json:"thumb"`
			DeleteURL string `json:"delete_url"`
			Id        string `json:"id"`
		} `json:"data"`
		Sucess bool `json:"success"`
		Status int  `json:"status"`
	}
	var ibb imageBBResponse
	err := json.NewDecoder(resp.Body).Decode(&ibb)
	if err != nil {
		return nil, err
	}

	if !ibb.Sucess {
		return nil, ErrImageUpload
	}

	return &models.Image{
		URL:       ibb.Data.URL,
		DeleteURL: ibb.Data.DeleteURL,
		Thumbnail: ibb.Data.Thumb.URL,
		Id:        ibb.Data.Id,
	}, nil
}
