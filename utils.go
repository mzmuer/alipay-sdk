package alipay

import (
	"bytes"
	"sort"
)

// 组成签名raw串
func GetSignatureContent(m map[string]string) string {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	// 对keys排序
	sort.Strings(keys)

	var buf bytes.Buffer
	for i, key := range keys {
		if m[key] == "" {
			continue
		}

		if i != 0 {
			buf.WriteString("&")
		}

		buf.WriteString(key)
		buf.WriteString("=")
		buf.WriteString(m[key])
	}

	return buf.String()
}
