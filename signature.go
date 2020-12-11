package crypto_rsa

// 使用RSAWithSHA1算法签名
func SignSha1WithRsa(data string, privateKey string) (string, error) {
	rsa := RSASecurity{}
	rsa.SetPrivateKey(privateKey)

	sign, err := rsa.SignSha1WithRsa(data)
	if err != nil {
		return "", err
	}

	return sign, err
}

// 使用RSAWithSHA256算法签名
func SignSha256WithRsa(data string, privateKey string) (string, error) {
	rsa := RSASecurity{}
	rsa.SetPrivateKey(privateKey)

	sign, err := rsa.SignSha256WithRsa(data)
	if err != nil {
		return "", err
	}
	return sign, err
}

// 使用RSAWithSHA1验证签名
func VerifySignSha1WithRsa(data string, signData string, publicKey string) error {
	rsa := RSASecurity{}
	rsa.SetPublicKey(publicKey)
	return rsa.VerifySignSha1WithRsa(data, signData)
}

// 使用RSAWithSHA256验证签名
func VerifySignSha256WithRsa(data string, signData string, publicKey string) error {
	rsa := RSASecurity{}
	rsa.SetPublicKey(publicKey)
	return rsa.VerifySignSha256WithRsa(data, signData)
}

