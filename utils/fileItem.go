package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type FileItem struct {
	FileName string    // 非必填，上传时会随机生成文件名
	MIMEType string    // 非必填，会根据文件内容自动识别类型
	Reader   io.Reader // Reader 和Content二选一必填
	Content  []byte    // Content 设置之后就不会使用Reader
}

// --
func (f *FileItem) GetMIMEType() string {
	if f.MIMEType == "" {
		f.MIMEType = _getMimeType(f.GetContent())
	}

	return f.MIMEType
}

func (f *FileItem) GetFileName() string {
	if f.FileName == "" {
		suffix := f.GetMIMEType()
		if s := strings.Split(suffix, "/"); len(s) == 2 {
			suffix = s[1]
		}

		return RandomString(8) + "." + suffix
	}

	return f.FileName
}

func (f *FileItem) GetContent() []byte {
	if len(f.Content) == 0 && f.Reader != nil {
		if b, err := ioutil.ReadAll(f.Reader); err != nil {
			fmt.Println("read file failed", err)
		} else {
			f.Content = b
		}
	}

	return f.Content
}

// ---
func _getMimeType(b []byte) string {
	suffix := _getFileSuffix(b)
	var mimeType string
	switch suffix {
	case "JPG":
		mimeType = "image/jpeg"
	case "GIF":
		mimeType = "image/gif"
	case "PNG":
		mimeType = "image/png"
	case "BMP":
		mimeType = "image/bmp"
	default:
		mimeType = "application/octet-stream"
	}

	return mimeType
}

func _getFileSuffix(b []byte) string {
	if b == nil || len(b) < 10 {
		return ""
	}

	if b[0] == 'G' && b[1] == 'I' && b[2] == 'F' {
		return "GIF"
	} else if b[1] == 'P' && b[2] == 'N' && b[3] == 'G' {
		return "PNG"
	} else if b[6] == 'J' && b[7] == 'F' && b[8] == 'I' && b[9] == 'F' {
		return "JPG"
	} else if b[0] == 'B' && b[1] == 'M' {
		return "BMP"
	}

	return ""

}
