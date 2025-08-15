package util

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSecret() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
