package main

import (
	"fmt"
	"gotaskapp/internal/app"
	"gotaskapp/internal/configs"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(`./config.yaml`)
	viper.AutomaticEnv() // for global env

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	fmt.Println("test")
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	fmt.Println(filepath.Join(basepath, "../"))
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
