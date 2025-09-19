package sec_test

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/digiconvent/d9t/utils/sec"
)

func TestRsaConversion(t *testing.T) {
	pk, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatal("don't expect this to fail")
	}

	pk2 := sec.StringToPrivateKey(sec.PrivateKeyToString(pk))

	if !pk2.Equal(pk) {
		t.Fatal("expected this to be identical")
	}
}
