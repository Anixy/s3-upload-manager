package util

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	// Mengubah byte menjadi string dengan base64 encoding
	randomString := base64.RawURLEncoding.EncodeToString(buffer)

	// Mengambil substring sepanjang length
	randomString = randomString[:length]

	return randomString, nil
}