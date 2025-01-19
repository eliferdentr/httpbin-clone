package interfaces

type RequestProcessorService interface {
	GetHeader(headers map[string][]string) map[string][]string
	GetMethod(method string) string
	GetBody(body interface{}) (interface{}, error)
	GetQueryParams(params map[string][]string) map[string]interface{}
	GetRequestDetails(method string, headers map[string][]string, params map[string][]string, body interface{}) map[string]interface{}
}