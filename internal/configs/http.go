package configs

import (
	"gotaskapp/internal/controllers"
	"gotaskapp/internal/routes"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

type ServerHttp struct {
	server  *gin.Engine
	gateway controllers.GatewayController
}

func (s *ServerHttp) HandleV1Route() {
	v1 := s.server.Group("/api/v1")

	routes.UserRoutes(v1, s.gateway)
}

func (s *ServerHttp) Start() error {
	s.HandleV1Route()
	return s.server.Run(viper.GetString(`app.port`))
}

func NewServiceHttp(controller controllers.GatewayController) *ServerHttp {
	return &ServerHttp{
		server:  gin.Default(),
		gateway: controller,
	}
}
