package crypto_rsa

import (
	"io/ioutil"
	"testing"
)

func TestSignSha1WithRsa(t *testing.T) {
	readData, err := ioutil.ReadFile("./rsa_private.pem")
	if err != nil {
		t.Fatal(err)
	}
	signature, err := SignSha1WithRsa("123", string(readData))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(signature)

	readData, err = ioutil.ReadFile("./rsa_public.pem")
	if err != nil {
		t.Fatal(err)
	}
	err = VerifySignSha1WithRsa("123", signature, string(readData))
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignSha256WithRsa(t *testing.T) {
	readData, err := ioutil.ReadFile("./rsa_private.pem")
	if err != nil {
		t.Fatal(err)
	}
	signature, err := SignSha256WithRsa("123", string(readData))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(signature)

	readData, err = ioutil.ReadFile("./rsa_public.pem")
	if err != nil {
		t.Fatal(err)
	}
	err = VerifySignSha256WithRsa("123", signature, string(readData))
	if err != nil {
		t.Fatal(err)
	}
}