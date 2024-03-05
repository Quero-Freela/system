package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(length int) (string, error) {
	bytes, err := GenerateRandomBytes(length)

	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

func GenerateRandomBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)

	_, err := rand.Read(bytes)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}
