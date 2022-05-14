package configs

import (
	"fmt"
	"gotaskapp/internal/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitConnectDB() {

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.username`)
	dbPass := viper.GetString(`database.password`)
	dbName := viper.GetString(`database.dbname`)
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", dbHost, dbUser, dbPass, dbName, dbPort)

	DB, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("database is ready.....")
	DB.AutoMigrate(&models.User{}) // use for example
}

func NewDatabase() *gorm.DB {
	return DB
}
