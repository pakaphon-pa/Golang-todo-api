package repository

func (s *TestSuite) Test_GetAll() {
	result, err := s.userRepo.Find()

	s.Assert().NoError(err)
	s.Assert().NotEmpty(result)
}
