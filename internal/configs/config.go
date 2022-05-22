package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var configuration Config

type Config struct {
	App      App      `mapstructure:"app"`
	Database Database `mapstructure:"database"`
	Redis    Redis    `mapstructure:"redis"`
	Jwt      Jwt      `mapstructure:"jwt"`
}

type App struct {
	Env   string `mapstructure:"env"`
	Port  string `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
}

type Redis struct {
	Dsn string `mapstructure:"dsn"`
}

type Jwt struct {
	Access  string `mapstructure:"access_secret"`
	Refresh string `mapstructure:"refresh_secret"`
}

func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.AutomaticEnv() // for global env

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	fmt.Println(configuration.App)
	fmt.Println(configuration.App.Debug)

	if configuration.App.Env == "development" {
		log.Println("Service RUN on Development mode")
	}
}

func GetConfigs() Config {
	return configuration
}
