package middlewarevalidator

import (
	"errors"
	"gotaskapp/internal/models"
	"gotaskapp/pkg/customError"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RoleValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var role models.RoleRequest
		if err := c.ShouldBindJSON(&role); err == nil {
			validate := validator.New()
			if err := validate.Struct(&role); err != nil {
				validationErrors := err.(validator.ValidationErrors)
				if errors.As(err, &validationErrors) {
					messages := make([]FieldsError, len(validationErrors))
					for i, fieldError := range validationErrors {
						messages[i] = FieldsError{fieldError.Field(), fieldError.Tag()}
					}
					c.Error(customError.NewHTTPError(
						http.StatusBadRequest,
						"validate",
						messages,
					))
				}
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
