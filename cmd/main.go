package main

import (
	"gotaskapp/internal/app"
	"gotaskapp/internal/configs"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(`./config.yaml`)
	viper.AutomaticEnv() // for global env

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var configuration configs.Config
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	if configuration.App.Env == "development" {
		log.Println("Service RUN on Development mode")
	}

	configs.InitRedis()
	configs.InitConnectDB()

	app, err := app.NewApplication().Application()
	if err != nil {
		panic(err)
	}

	app.Start()
}
