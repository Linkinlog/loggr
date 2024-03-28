package assets

import (
	"embed"
	"io/fs"
	"path/filepath"
)

//go:embed *
var assets embed.FS

func NewAssets() Assets {
	return Assets{fs: files()}
}

type Assets struct {
	fs embed.FS
}

func (a Assets) Open(name string) (fs.File, error) {
	file, err := a.fs.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		index := filepath.Join(name, "index.html")
		if _, err := a.fs.Open(index); err != nil {
			closeErr := file.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}
	return file, nil
}

func files() embed.FS {
	return assets
}
