package utils

import (
	"io/ioutil"
	"os"
)

type FileItem struct {
	FileName string
	MimeType string
	Content  []byte
}

// --
func NewFileItem(filePath string) (*FileItem, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return &FileItem{
		FileName: f.Name(),
		MimeType: "",
		Content:  b,
	}, nil
}

// --- TODO:
