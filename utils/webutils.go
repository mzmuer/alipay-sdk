package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"

	"github.com/mzmuer/alipay-sdk/constant"
)

func DoPost(postUrl string, m map[string]string) ([]byte, error) {
	var (
		cType = "application/x-www-form-urlencoded;charset=" + m[constant.Charset]
		query = _buildQuery(m)
	)

	resp, err := http.Post(postUrl, cType, strings.NewReader(query.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// --
var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func DoPostUploadFile(postUrl string, m map[string]string, fileParams map[string]*FileItem) ([]byte, error) {
	if fileParams == nil || len(fileParams) == 0 {
		return DoPost(postUrl, m)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for fieldName, file := range fileParams {
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition",
			fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
				quoteEscaper.Replace(fieldName), quoteEscaper.Replace(file.GetFileName())))
		h.Set("Content-Type", file.GetMimeType())

		part, err := writer.CreatePart(h)
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(part, bytes.NewReader(file.GetContent()))
		if err != nil {
			return nil, err
		}
	}

	for key, val := range m {
		_ = writer.WriteField(key, val)
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", postUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType()+" ;charset="+m[constant.Charset])
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func _buildQuery(params map[string]string) url.Values {
	query := url.Values{}
	for key, val := range params {
		// 忽略参数名或参数值为空的参数
		if key == "" || val == "" {
			continue
		}

		query.Set(key, val)
	}

	return query
}
