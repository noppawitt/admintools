package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// Encrypt returns encrypted data in base16
func Encrypt(data string, key string) (string, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	byteArray := gcm.Seal(nonce, nonce, []byte(data), nil)
	return fmt.Sprintf("%x", byteArray), err
}
