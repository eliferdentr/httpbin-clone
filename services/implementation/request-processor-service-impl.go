package services

import (
	"errors"
	"fmt"
	"strings"
)

type RequestProcessorServiceImpl struct{}

func NewRequestProcessorServiceImpl() *RequestProcessorServiceImpl {
	return &RequestProcessorServiceImpl{}
}

func (r *RequestProcessorServiceImpl) GetHeader(headers map[string][]string) map[string][]string {
	//filter the headers or check if a header exists

	filteredHeaders := make(map[string][]string)

	for key, values := range headers {
		nonEmptyValues := []string{}
		for _, value := range values {
			trimmedValue := strings.TrimSpace(value) // Boşlukları temizle
			if trimmedValue != "" {
				nonEmptyValues = append(nonEmptyValues, trimmedValue)
			}
		}
		if len(nonEmptyValues) > 0 {
			filteredHeaders[key] = nonEmptyValues
		}
	}
	return filteredHeaders
}
func (r *RequestProcessorServiceImpl) GetMethod(method string) string {
	return strings.TrimSpace(method)

}
func (r *RequestProcessorServiceImpl) GetBody(body interface{}) (interface{}, error) {
	parsedBody, ok := body.(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid body format")
	}
	return parsedBody, nil

}
func (r *RequestProcessorServiceImpl) GetQueryParams(params map[string][]string) map[string][]string {
	//return the query params in a certain format
}
func (r *RequestProcessorServiceImpl) GetRequestDetails(method string, headers map[string][]string, params map[string][]string, body interface{}) map[string]interface{} {
	// return details of the request

}
