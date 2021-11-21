package errmsg

import "encoding/json"

type ApiError struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (error ApiError) Error() string {
	bytes, err := json.Marshal(error)
	if err != nil {
		return "{\"code\": \"SystemError\", \"message\": \"parse error wrong!\"}"
	}
	return string(bytes)
}
