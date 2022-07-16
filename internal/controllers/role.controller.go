package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
}

func (r *RoleController) CreateRole(g *gin.Context) {

	g.String(http.StatusCreated, "success")
}

func NewRoleController() *RoleController {
	return &RoleController{}
}
