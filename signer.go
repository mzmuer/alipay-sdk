package alipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

type signer struct {
	PrivateKey string
}

func (s *signer) Sign(sourceContent string, signType string, charset string) (string, error) {
	p, _ := pem.Decode([]byte(s.PrivateKey))
	key, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		return "", err
	}

	hashed := sha256.Sum256([]byte(sourceContent))

	var signed []byte
	if signType == SignTypeRSA2 {
		signed, err = rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed[:])
		if err != nil {
			return "", err
		}
	} else if signType == SignTypeRSA {
		signed, err = rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA1, hashed[:])
		if err != nil {
			return "", err
		}
	} else {
		return "", fmt.Errorf("unknown sign_type[%s]", signType)
	}

	return base64.StdEncoding.EncodeToString(signed), nil
}
