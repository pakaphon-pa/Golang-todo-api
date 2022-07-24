package main

import (
	"gotaskapp/internal/app"
	"gotaskapp/internal/configs"
)

func main() {
	configs.LoadConfig()
	config := configs.GetConfigs()
	configs.InitRedis(config)
	configs.InitConnectDB(config)

	app, err := app.NewApplication().Application()
	if err != nil {
		panic(err)
	}

	app.Start()
}
