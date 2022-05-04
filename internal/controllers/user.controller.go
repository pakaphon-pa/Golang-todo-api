package controllers

import (
	"gotaskapp/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService models.UserServiceInterface
}

func (c *UserController) GetUser(g *gin.Context) {
	result, _ := c.userService.Get()

	g.JSON(http.StatusOK, result)
}

func NewUserController(user models.UserServiceInterface) *UserController {
	return &UserController{
		userService: user,
	}
}
