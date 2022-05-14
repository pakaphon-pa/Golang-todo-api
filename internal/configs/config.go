package configs

type Config struct {
	App App `mapstructure:"app"`
}

type App struct {
	Env   string `mapstructure:"env"`
	Port  string `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
}

// func LoadConfig(path string) {
// 	viper.SetConfigName("config.yaml")
// 	viper.AddConfigPath("../../../")
// 	viper.AutomaticEnv() // for global env

// 	if err := viper.ReadInConfig(); err != nil {
// 		panic(err)
// 	}

// 	var configuration Config
// 	err := viper.Unmarshal(&configuration)
// 	if err != nil {
// 		log.Fatalf("unable to decode into struct, %v", err)
// 	}
// 	fmt.Println(configuration.App)
// 	fmt.Println(configuration.App.Debug)

// 	if configuration.App.Env == "development" {
// 		log.Println("Service RUN on Development mode")
// 	}
// }
