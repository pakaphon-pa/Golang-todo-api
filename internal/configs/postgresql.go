package configs

import (
	"fmt"
	"gotaskapp/internal/models"
	"gotaskapp/pkg/utility"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitConnectDB(config Config) {

	fmt.Println(config.Database)

	dbHost := config.Database.Host
	dbPort := config.Database.Port
	dbUser := config.Database.Username
	dbPass := config.Database.Password
	dbName := config.Database.DbName
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", dbHost, dbUser, dbPass, dbName, dbPort)

	DB, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	utility.LoadSampleData(DB)

	fmt.Println("database is ready.....")
	DB.AutoMigrate(&models.User{}, &models.Role{}) // use for example

}

func NewDatabase() *gorm.DB {
	return DB
}
