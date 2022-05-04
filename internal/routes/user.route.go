package routes

import (
	"gotaskapp/internal/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(v1 *gin.RouterGroup, c controllers.GatewayController) {
	user := v1.Group("/users")
	user.GET("", c.UserController.GetUser)
}
