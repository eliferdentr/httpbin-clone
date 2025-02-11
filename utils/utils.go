package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"
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

func ExtractKeyValue (pair string) (string, string) {

}

func SplitByCommas (header string) []string {
	var result []string
	var current string
	inQuotes := false

	for _,char := range header {
		switch char {
		case '"':
			inQuotes = !inQuotes
		case ',':
			if !inQuotes {
				result = append(result, strings.TrimSpace(current))
				current = ""
				continue
			}
		}
		current += string(char)
	}
	if current != "" {
		result = append(result, strings.TrimSpace(current))
	}
	return result
}
