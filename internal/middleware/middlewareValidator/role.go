package middlewarevalidator

import (
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
				c.Error(customError.NewHTTPError(
					http.StatusBadRequest,
					"validate",
					err.Error(),
				))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
