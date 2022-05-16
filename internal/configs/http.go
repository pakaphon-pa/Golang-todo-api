package configs

import (
	"gotaskapp/internal/controllers"
	"gotaskapp/internal/routes"

	"github.com/gin-gonic/gin"
)

type ServerHttp struct {
	server  *gin.Engine
	gateway controllers.GatewayController
	configs Config
}

func (s *ServerHttp) HandleV1Route() {
	v1 := s.server.Group("/api/v1")

	routes.UserRoutes(v1, s.gateway)
}

func (s *ServerHttp) Start() error {
	s.HandleV1Route()
	return s.server.Run(s.configs.App.Port)
}

func NewServiceHttp(controller controllers.GatewayController, configs Config) *ServerHttp {
	return &ServerHttp{
		server:  gin.Default(),
		gateway: controller,
		configs: configs,
	}
}
