package utils

import (
	"io"
)

type FileItem struct {
	FileName string
	MimeType string
	Content  io.Reader
}

// --
func NewFileItem(fileName string, r io.Reader, mimeType string) *FileItem {
	return &FileItem{
		FileName: fileName,
		MimeType: mimeType,
		Content:  r,
	}
}

// --- TODO:
