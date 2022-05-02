package main

import (
	"gotaskapp/internal/app"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(`./config.yaml`)
	viper.AutomaticEnv() // for global env

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetString(`app.env`) == "development" {
		log.Println("Service RUN on Development mode")
	}

	app, err := app.NewApplication().Application()
	if err != nil {
		panic(err)
	}

	app.Start()
}
