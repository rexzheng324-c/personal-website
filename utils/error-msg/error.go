package error_msg

import "encoding/json"

type ApiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewApiError(code string, message string) *ApiError {
	return &ApiError{
		Code: code,
		Message: message,
	}
}

func (error *ApiError) Error() string {
	bytes, err := json.Marshal(error)
	if err != nil {
		return "{\"code\": \"500\", \"message\": \"parse error wrong!\"}"
	}
	return string(bytes)
}
