package rsa

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io/ioutil"
)

type RSASecurity struct {
	pubStr string          //公钥字符串
	priStr string          //私钥字符串
	pubkey *rsa.PublicKey  //公钥
	prikey *rsa.PrivateKey //私钥
}

// 设置公钥
func (rsas *RSASecurity) SetPublicKey(pubStr string) (err error) {
	rsas.pubStr = pubStr
	rsas.pubkey, err = rsas.GetPublickey()
	return err
}

// 设置私钥
func (rsas *RSASecurity) SetPrivateKey(priStr string) (err error) {
	rsas.priStr = priStr
	rsas.prikey, err = rsas.GetPrivatekey()
	return err
}

// *rsa.PublicKey
func (rsas *RSASecurity) GetPrivatekey() (*rsa.PrivateKey, error) {
	return byte2PriKey([]byte(rsas.priStr))
}

// *rsa.PrivateKey
func (rsas *RSASecurity) GetPublickey() (*rsa.PublicKey, error) {
	return byte2PubKey([]byte(rsas.pubStr))
}

// 公钥加密
func (rsas *RSASecurity) PubKeyENCTYPT(input []byte) ([]byte, error) {
	if rsas.pubkey == nil {
		return []byte(""), errors.New(`Please set the public key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := pubKeyIO(rsas.pubkey, bytes.NewReader(input), output, true)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(output)
}

// 公钥解密
func (rsas *RSASecurity) PubKeyDECRYPT(input []byte) ([]byte, error) {
	if rsas.pubkey == nil {
		return []byte(""), errors.New(`Please set the public key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := pubKeyIO(rsas.pubkey, bytes.NewReader(input), output, false)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(output)
}

// 私钥加密
func (rsas *RSASecurity) PriKeyENCTYPT(input []byte) ([]byte, error) {
	if rsas.prikey == nil {
		return []byte(""), errors.New(`Please set the private key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := priKeyIO(rsas.prikey, bytes.NewReader(input), output, true)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(output)
}

// 私钥解密
func (rsas *RSASecurity) PriKeyDECRYPT(input []byte) ([]byte, error) {
	if rsas.prikey == nil {
		return []byte(""), errors.New(`Please set the private key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := priKeyIO(rsas.prikey, bytes.NewReader(input), output, false)
	if err != nil {
		return []byte(""), err
	}

	return ioutil.ReadAll(output)
}

// SHA1签名
func (rsas *RSASecurity) SignSha1WithRsa(data string) (string, error) {
	sha1Hash := sha1.New()
	s_data := []byte(data)
	sha1Hash.Write(s_data)
	hashed := sha1Hash.Sum(nil)

	signByte, err := rsa.SignPKCS1v15(rand.Reader, rsas.prikey, crypto.SHA1, hashed)
	sign := base64.StdEncoding.EncodeToString(signByte)
	return string(sign), err
}

// SHA256签名
func (rsas *RSASecurity) SignSha256WithRsa(data string) (string, error) {
	sha256Hash := sha256.New()
	s_data := []byte(data)
	sha256Hash.Write(s_data)
	hashed := sha256Hash.Sum(nil)

	signByte, err := rsa.SignPKCS1v15(rand.Reader, rsas.prikey, crypto.SHA256, hashed)
	sign := base64.StdEncoding.EncodeToString(signByte)
	return string(sign), err
}

// SHA1验签
func (rsas *RSASecurity) VerifySignSha1WithRsa(data string, signData string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	hash := sha1.New()
	hash.Write([]byte(data))
	return rsa.VerifyPKCS1v15(rsas.pubkey, crypto.SHA1, hash.Sum(nil), sign)
}

// SHA256验签
func (rsas *RSASecurity) VerifySignSha256WithRsa(data string, signData string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	hash := sha256.New()
	hash.Write([]byte(data))

	return rsa.VerifyPKCS1v15(rsas.pubkey, crypto.SHA256, hash.Sum(nil), sign)
}
