package test

import (
	"context"
	"fmt"
	"gotaskapp/internal/configs"
	"gotaskapp/internal/models"
	"gotaskapp/internal/repository"
	"gotaskapp/internal/services"
	"log"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	container "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	driverDb "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestSuite struct {
	suite.Suite
	db *gorm.DB

	userRepo    models.UserRepositoryInterface
	userService models.UserServiceInterface
}

func TestInit(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupSuite() {
	log.Println("Prepare DB....")
	viper.SetConfigFile(`../../config.yaml`)
	configs.LoadConfig("../../config.yaml")

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

	dbHost := "localhost"
	dbName := viper.GetString(`database.dbname`)

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", dbHost, dbUser, dbPass, dbName, port.Port())

	DB, err := gorm.Open(driverDb.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("database is ready.....")
	DB.AutoMigrate(&models.User{})

	s.db = DB

}

func (s *TestSuite) SetupTest() {
	s.userRepo = repository.NewUserRepository(s.db)
	s.userService = services.NewUserService(s.userRepo)
}
