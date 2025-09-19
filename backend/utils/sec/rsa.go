package sec

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func PrivateKeyToString(key *rsa.PrivateKey) string {
	privBytes := x509.MarshalPKCS1PrivateKey(key)

	privPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	})

	return string(privPem)
}

func StringToPrivateKey(key string) *rsa.PrivateKey {
	if key == "" {
		return nil
	}
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil
	}

	privKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	return privKey
}
