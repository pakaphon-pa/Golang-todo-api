package controllers

import (
	"gotaskapp/internal/models"
	"gotaskapp/pkg/customError"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService models.UserServiceInterface
}

func (c *UserController) GetUser(g *gin.Context) {
	// result, _ := c.userService.Get()

	// g.JSON(http.StatusOK, result)
	g.Error(customError.NewHTTPError(404, "Not found", "Not found"))
}

func (c *UserController) GetUserById(g *gin.Context) {
	_, err := c.userService.GetById()

	if err != nil {
		g.JSON(http.StatusOK, err)
	}
}

func NewUserController(user models.UserServiceInterface) *UserController {
	return &UserController{
		userService: user,
	}
}
