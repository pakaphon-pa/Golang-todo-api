package test

import "gotaskapp/internal/models"

func (s *TestSuite) Test_CreateRole() {
	var role = models.Role{
		Name: "admin",
	}
	result, err := s.roleRepo.Create(&role)

	s.Assert().NoError(err)
	s.Assert().NotNil(result)
}

func (s *TestSuite) Test_GetAllRole() {
	result, err := s.roleRepo.Get()

	s.Assert().NoError(err)
	s.Assert().Len(result, 1)
}
