package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetKeyValueMap(input map[string][]string) map[string]string {
	args := map[string]string{}
	for k,v := range input {
		if len(v) > 0 {
			args[k] = v[0]
		}
	}
	return args
}

func GetJSONBody (c *gin.Context)( string, map[string]interface{}) {
	var rawBody string
	var jsonBody map[string]interface{}

	if c.Request.Body != nil {
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		rawBody = string(bodyBytes)
		_ = json.Unmarshal((bodyBytes), &jsonBody)
		// Body tekrar kullanılabilsin diye resetle

		c.Request.Body = io.NopCloser(strings.NewReader(rawBody))
	}
	return rawBody, jsonBody
}


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



func BuildResponse ( context *gin.Context, requestBody any) gin.H {
	args := context.Request.URL.Query()
	headers := make(map[string]string)

	for k, v := range context.Request.Header {
		headers [k] = strings.Join(v, ", ")
	}

	origin := context.ClientIP()

	scheme := "http"

	if context.Request.TLS != nil {
		scheme = "https"
	}

	url := fmt.Sprintf(
		"%s://%s%s", scheme, context.Request.Host, context.Request.RequestURI)

	response := gin.H{
		"args" : args,
		"headers": headers,
		"origin": origin,
		"url": url,
	}

	if requestBody != nil {
		response ["data"] = requestBody
	}

	return response
}

func GetNParam(c *gin.Context) (int, bool) { 
	nParamStr := c.Param("n") 
	nParam, err := strconv.Atoi(nParamStr) 
	if err != nil || nParam <= 0 { 
		return 0, true 
	} 
	return nParam, false 
}