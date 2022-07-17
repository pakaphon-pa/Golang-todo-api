package customError

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpError struct {
	Code    int
	Key     string      `json:"error"`
	Message interface{} `json:"message"`
}

func NewHTTPError(code int, key string, msg interface{}) *HttpError {
	return &HttpError{
		Code:    code,
		Key:     key,
		Message: msg,
	}
}

func HttpErrorCustomMiddleware() gin.HandlerFunc {
	return httpErrorCustom(gin.ErrorTypeAny)
}

func httpErrorCustom(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		log.Println("Handle APP error")
		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var parsedError *HttpError
			switch err := err.(type) {
			case *HttpError:
				parsedError = err
			default:
				parsedError = &HttpError{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}
			c.IndentedJSON(parsedError.Code, parsedError)
			c.Abort()
			return
		}

	}
}

// Error makes it compatible with `error` interface.
func (e *HttpError) Error() string {
	return fmt.Sprintf("code=%d, message=%v, message=%v", e.Code, e.Message, e.Key)
}
