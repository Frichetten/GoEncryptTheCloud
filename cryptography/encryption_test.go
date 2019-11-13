package cryptography

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	// Encrypt
	plaintext := []byte("Hello Friend!")
	key := SHA256Hash("password")

	output, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatal(err)
	}

	// Decrypt
	newKey := SHA256Hash("password")
	newPlaintext, err := Decrypt(output, newKey)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(plaintext, newPlaintext) {
		t.Fatal("Plaintexts are not the same")
	}
}

func TestSHA256Hash(t *testing.T) {
	key := SHA256Hash("friend")
	// String is sha256 hash of friend
	if hex.EncodeToString(key) != "cde48537ca2c28084ff560826d0e6388b7c57a51497a6cb56f397289e52ff41b" {
		t.Fatal("Hash does not match")
	}
}
