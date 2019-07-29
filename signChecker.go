package alipay

type signChecker struct {
	publicKey string
}

func (s *signer) Check(sourceContent string, signature string, signType string, charset string) (bool, error) {
	return false, nil
}
