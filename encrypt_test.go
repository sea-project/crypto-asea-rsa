package crypto_rsa

import (
	"io/ioutil"
	"testing"
)

// 公钥加密私钥解密
func TestPublicEncrypt(t *testing.T) {
	readData, err := ioutil.ReadFile("./rsa_public.pem")
	data, err := PublicEncrypt("123", string(readData))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)


	readData, err = ioutil.ReadFile("./rsa_private.pem")
	str, err := PriKeyDecrypt(data, string(readData))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}

// 公钥加密私钥解密
func TestPriKeyEncrypt(t *testing.T) {
	readData, err := ioutil.ReadFile("./rsa_private.pem")
	data, err := PriKeyEncrypt("123", string(readData))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	readData, err = ioutil.ReadFile("./rsa_public.pem")
	str, err := PublicDecrypt(data, string(readData))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}
