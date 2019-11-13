package cryptography

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	plaintext := []byte("Hello Friend!")
	key := &[32]byte{}
	for i := 0; i < 32; i++ {
		key[i] = '0'
	}

	_, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatal(err)
	}
}
