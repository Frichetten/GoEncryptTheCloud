package cryptography

import (
	"bytes"
	"testing"
)

func TestEncrypt(t *testing.T) {
	// Encrypt
	plaintext := []byte("Hello Friend!")
	key := SHA256Hash("password")

	output, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatal(err)
	}

	// Decrypt
	newPlaintext, err := Decrypt(output, key)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(plaintext, newPlaintext) {
		t.Fatal("Plaintexts are not the same")
	}
}
