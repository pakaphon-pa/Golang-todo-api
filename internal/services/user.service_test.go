package services

func (s *ServiceTestSuite) Test_GetAll() {
	result, err := s.userService.Get()

	s.Assert().NoError(err)
	s.Assert().NotEmpty(result)
}

func (s *ServiceTestSuite) Test_GetById() {
	_, err := s.userService.GetById()

	s.Assert().Error(err)
}
