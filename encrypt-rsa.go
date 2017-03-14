package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"time"
	"fmt"
)

func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(rsaPublicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(rsaPrivateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func main() {
	t0 := time.Now()
	data, err := RsaEncrypt([]byte("wangqiang11@le.com"))
	t1 := time.Now()
	fmt.Printf("encrypting cost : %v\n", t1.Sub(t0))
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted : " + string(data))

	t0 = time.Now()
	originData, err := RsaDecrypt(data)
	t1 = time.Now()
	fmt.Printf("decrypting cost : %v\n", t1.Sub(t0))
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted : " + string(originData))
}
