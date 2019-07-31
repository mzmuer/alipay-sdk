package alipay

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

type signChecker struct {
	publicKey string
}

func (s *signChecker) Check(sourceContent string, signature string, signType string, charset string) (bool, error) {
	pubKey, err := genPubKey([]byte(s.publicKey))
	if err != nil {
		return false, err
	}

	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}

	h := sha256.New()
	h.Write([]byte(sourceContent))

	if signType == SignTypeRSA2 {
		err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, h.Sum(nil), decoded)
	} else if signType == SignTypeRSA {
		err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA1, h.Sum(nil), decoded)
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func genPubKey(key []byte) (*rsa.PublicKey, error) {
	encodedKey, err := base64.StdEncoding.DecodeString(string(key))
	if err != nil {
		return nil, err
	}

	pkix, err := x509.ParsePKIXPublicKey(encodedKey)
	if err != nil {
		return nil, fmt.Errorf("unable to parse pxix key")
	}

	pubKey, ok := pkix.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("aliPubKey can not be parsed to rsa.PublicKey")
	}

	return pubKey, nil
}
