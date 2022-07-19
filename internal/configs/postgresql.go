package configs

import (
	"fmt"
	"gotaskapp/internal/models"
	"log"
	"os"

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

	fmt.Println("database is ready.....")
	DB.AutoMigrate(&models.User{}, &models.Role{}) // use for example

	pathToFile := "../../fixtures.sql"
	q, err := os.ReadFile(pathToFile)
	if err != nil {
		log.Fatal("fixtures:", err)
	}
	DB.Exec(string(q))

}

func NewDatabase() *gorm.DB {
	return DB
}
