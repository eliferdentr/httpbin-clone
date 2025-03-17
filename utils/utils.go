package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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



func SplitByCommas(header string) []string {
	var result []string
	var current string
	inQuotes := false

	for _, char := range header {
		switch char {
		case '"':
			inQuotes = !inQuotes // Tırnak açıp kapatma durumunu takip et
		case ',':
			if !inQuotes {
				// Eğer tırnak içinde değilsek, virgülde böl
				result = append(result, strings.TrimSpace(current))
				current = ""
				continue
			}
		}
		current += string(char)
	}

	// Son parçayı ekleyelim
	if current != "" {
		result = append(result, strings.TrimSpace(current))
	}

	return result
}
func ExtractKeyValue(pair string) (string, string) {
	parts := strings.SplitN(pair, "=", 2)
	if len(parts) != 2 {
		return "", ""
	}

	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])

	// Eğer value çift tırnak içindeyse, tırnakları temizle
	if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
		value = value[1 : len(value)-1]
	}

	return key, value
}

func GetPositiveIntParam(context *gin.Context, paramName string) (int, error) {
	paramStr := context.Param(paramName) // Parametreyi al
	param, err := strconv.Atoi(paramStr) // String to int dönüşümü
	if err != nil || param <= 0 {
		return 0, fmt.Errorf("Invalid parameter. Please provide a positive integer.")
	}
	return param, nil
}