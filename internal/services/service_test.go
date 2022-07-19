package services

import (
	"gotaskapp/internal/models"
	"gotaskapp/internal/repository"
	testcontainer "gotaskapp/pkg/testContainer"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ServiceTestSuite struct {
	suite.Suite
	db *gorm.DB

	userService models.UserServiceInterface
}

func TestServiceTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) SetupSuite() {
	DB := testcontainer.PostgresqlTestContainer()
	s.db = DB
}

func (s *ServiceTestSuite) SetupTest() {
	userRepo := repository.NewUserRepository(s.db)
	s.userService = NewUserService(userRepo)
}
