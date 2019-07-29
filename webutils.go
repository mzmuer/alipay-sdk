package alipay

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
