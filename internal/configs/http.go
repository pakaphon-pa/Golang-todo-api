package configs

import (
	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

type ServerHttp struct {
	server *gin.Engine
}

func (s *ServerHttp) Start() error {

	return s.server.Run(viper.GetString(`app.port`))
}

func NewServiceHttp() *ServerHttp {
	return &ServerHttp{
		server: gin.Default(),
	}
}
