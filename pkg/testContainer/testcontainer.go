package testcontainer

import (
	"context"
	"fmt"
	"gotaskapp/internal/configs"
	"gotaskapp/internal/models"
	"gotaskapp/pkg/utility"

	"log"

	"github.com/docker/go-connections/nat"
	"github.com/spf13/viper"
	container "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	driverDb "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresqlTestContainer() *gorm.DB {

	log.Println("Prepare DB....")
	configs.LoadConfig()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.username`)
	dbPass := viper.GetString(`database.password`)
	postgresPort := nat.Port(dbPort)
	postgres, err := container.GenericContainer(context.Background(),
		container.GenericContainerRequest{
			ContainerRequest: container.ContainerRequest{
				Image:        "postgres",
				ExposedPorts: []string{postgresPort.Port()},
				Env: map[string]string{
					"POSTGRES_PASSWORD": dbPass,
					"POSTGRES_USER":     dbUser,
				},
				WaitingFor: wait.ForAll(
					wait.ForLog("database system is ready to accept connections"),
					wait.ForListeningPort(postgresPort),
				),
			},
			Started: true, // auto-start the container
		})

	if err != nil {
		log.Fatal("start:", err)
	}

	port, err := postgres.MappedPort(context.Background(), postgresPort)
	if err != nil {
		log.Fatal("map:", err)
	}

	dbHost, err := postgres.Host(context.Background())
	if err != nil {
		log.Fatal("host:", err)
	}
	dbName := viper.GetString(`database.dbname`)

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", dbHost, dbUser, dbPass, dbName, port.Port())

	DB, err := gorm.Open(driverDb.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("database is ready.....")
	DB.AutoMigrate(&models.User{})
	utility.LoadSampleData(DB)

	return DB
}
