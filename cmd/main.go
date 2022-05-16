package main

import (
	"gotaskapp/internal/app"
	"gotaskapp/internal/configs"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(`./config.yaml`)
	viper.AutomaticEnv() // for global env

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	configs.LoadConfig("../../config.yaml")

	config := configs.GetConfigs()
	configs.InitRedis(config)
	configs.InitConnectDB(config)

	app, err := app.NewApplication().Application()
	if err != nil {
		panic(err)
	}

	app.Start()
}
