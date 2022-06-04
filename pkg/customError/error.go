package customError

import "fmt"

type HttpError struct {
	Code    int
	Key     string `json:"error"`
	Message string `json:"message"`
}

func NewHTTPError(code int, key string, msg string) *HttpError {
	return &HttpError{
		Code:    code,
		Key:     key,
		Message: msg,
	}
}

// Error makes it compatible with `error` interface.
func (e *HttpError) Error() string {
	return fmt.Sprintf("code=%d, message=%v, message=%v", e.Code, e.Message, e.Key)
}
