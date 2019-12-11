package signature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/mzmuer/alipay-sdk"
)

type Signer struct {
	privateKey *rsa.PrivateKey
}

func NewSigner(privateKey []byte) (*Signer, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}

	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: " + err.Error())
	}

	return &Signer{privateKey: priKey}, err
}

func (s *Signer) Sign(sourceContent string, signType string, charset string) (string, error) {
	var (
		hashed = sha256.Sum256([]byte(sourceContent))
		signed []byte
		err    error
	)

	if signType == alipay.SignTypeRSA2 {
		signed, err = rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, hashed[:])
		if err != nil {
			return "", err
		}
	} else if signType == alipay.SignTypeRSA {
		signed, err = rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA1, hashed[:])
		if err != nil {
			return "", err
		}
	} else {
		return "", fmt.Errorf("unknown sign_type[%s]", signType)
	}

	return base64.StdEncoding.EncodeToString(signed), nil
}
