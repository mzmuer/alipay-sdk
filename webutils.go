package alipay

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/mzmuer/alipay-sdk/utils"
)

func doPost(postUrl string, m map[string]string) ([]byte, error) {
	var (
		cType = "application/x-www-form-urlencoded;charset=" + m[Charset]
		query = _buildQuery(m)
	)

	resp, err := http.Post(postUrl, cType, strings.NewReader(query.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func doPostUploadFile(postUrl string, m map[string]string, fileParams map[string]*utils.FileItem) ([]byte, error) {
	if fileParams == nil || len(fileParams) == 0 {
		return doPost(postUrl, m)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for fieldName, file := range fileParams {
		part, err := writer.CreateFormFile(fieldName, filepath.Base(file.FileName))
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(part, file.Content)
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

	req.Header.Set("Content-Type", writer.FormDataContentType()+" ;charset="+m[Charset])
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
