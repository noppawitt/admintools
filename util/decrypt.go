package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

// Decrypt return decrypted data from base16
func Decrypt(encryptedData string, key string) (string, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	bytesData, err := hex.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(bytesData) < nonceSize {
		return "", errors.New("encrypted data too short")
	}

	nonce, bytesData := bytesData[:nonceSize], bytesData[nonceSize:]
	bytesData, err = gcm.Open(nil, nonce, bytesData, nil)
	if err != nil {
		return "", err
	}
	return string(bytesData), err
}
