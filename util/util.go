package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// ToSnakeCase converts camelCase to snake_case
func ToSnakeCase(str string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// HashPassword returns hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

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
