package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateNonce(size int) (string, error) {
	// Generate random byte
	nonceBytes := make([]byte, size)
	_, err := rand.Read(nonceBytes)
	if err != nil {
		return "", err
	}

	// SHA-256
	hash := sha256.Sum256(nonceBytes)

	// return string in Hex format
	return hex.EncodeToString(hash[:]), nil
}