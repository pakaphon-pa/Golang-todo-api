package repository

import (
	"gotaskapp/internal/models"
	testcontainer "gotaskapp/pkg/testContainer"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type TestSuite struct {
	suite.Suite
	db *gorm.DB

	userRepo models.UserRepositoryInterface
	roleRepo models.RoleRepositoryInterface
}

func TestInit(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupSuite() {
	DB := testcontainer.PostgresqlTestContainer()

	s.db = DB
}

func (s *TestSuite) SetupTest() {
	s.userRepo = NewUserRepository(s.db)
	s.roleRepo = NewRoleRepository(s.db)
}
