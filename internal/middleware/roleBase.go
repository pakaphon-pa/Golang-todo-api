package middleware

import (
	"gotaskapp/pkg/customError"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorization(authRoles []string) gin.HandlerFunc {
	return func(context *gin.Context) {

		if len(context.Keys) == 0 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, customError.HttpError{
				Code:    http.StatusUnauthorized,
				Key:     "Unauthorized",
				Message: "Unauthorized access",
			})
		}

		roles := context.Keys["Roles"]

		if roles == nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, customError.HttpError{
				Code:    http.StatusUnauthorized,
				Key:     "Unauthorized",
				Message: "Unauthorized access",
			})
		}

		rolesList := roles.([]string)
		validation := make(map[string]int)
		for _, val := range rolesList {
			validation[val] = 0
		}

		for _, val := range authRoles {
			if _, ok := validation[val]; !ok {
				context.AbortWithStatusJSON(http.StatusUnauthorized, customError.HttpError{
					Code:    http.StatusUnauthorized,
					Key:     "Unauthorized",
					Message: "Unauthorized access",
				})
			}
		}

	}
}
