package services

import (
	"errors"
	"strings"
	"time"
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
func (r *RequestProcessorServiceImpl) GetQueryParams(params map[string][]string) map[string]interface{}{
	//return the query params in json format
	formattedParams := make(map[string]interface{})
	for key, values := range params {
		if len(values) == 1 {
			formattedParams[key] = values[0]
		} else if len(values) > 1 {
			formattedParams[key] = values
		} else {
			formattedParams[key] = nil
		}
	}

	return formattedParams
}
func (r *RequestProcessorServiceImpl) GetRequestDetails(method string, headers map[string][]string, params map[string][]string, body interface{}) map[string]interface{} {
	queryParams := r.GetQueryParams(params)
	filteredHeaders := r.GetHeader(headers)
	processedBody, _ := r.GetBody(body)

	totalQueryParams := len(queryParams)
	totalHeaders := len(headers)

	return map[string] interface {} {
		"method" : method,
		"valid_headers" : filteredHeaders,
		"total_valid_headers" : totalHeaders,
		"valid_query_params" : queryParams,
		"total_valid_query_keys" : totalQueryParams,
		"body" : processedBody,
		"meta" : map [string] interface{} {
			"processed_at" : time.Now(),
		},
	}


}
