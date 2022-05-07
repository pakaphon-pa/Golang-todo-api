package app

import (
	"fmt"
	"gotaskapp/internal/configs"
	"gotaskapp/internal/controllers"
	"gotaskapp/internal/repository"
	"gotaskapp/internal/services"

	"go.uber.org/dig"
)

type application struct {
	container *dig.Container
}

func (a *application) Application() (*configs.ServerHttp, error) {
	fmt.Println("Function use for set dependency injection by uber-dig")
	appConstructors := []interface{}{
		configs.NewServiceHttp,
		configs.NewDatabase,
		configs.GetRedis,
		repository.NewUserRepository,
		services.NewUserService,
		controllers.NewUserController,
	}

	for _, app := range appConstructors {
		if err := a.container.Provide(app); err != nil {
			return nil, err
		}
	}

	var result *configs.ServerHttp
	err := a.container.Invoke(func(a *configs.ServerHttp) {
		result = a
	})

	return result, err
}

func NewApplication() *application {
	return &application{
		container: dig.New(),
	}
}
