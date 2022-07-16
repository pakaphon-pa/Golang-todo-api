package routes

import (
	"gotaskapp/internal/controllers"
	middlewarevalidator "gotaskapp/internal/middleware/middlewareValidator"

	"github.com/gin-gonic/gin"
)

func RoleRoutes(v1 *gin.RouterGroup, c controllers.GatewayController) {
	user := v1.Group("/role")
	user.POST("", middlewarevalidator.RoleValidator(), c.RoleController.CreateRole)
}
