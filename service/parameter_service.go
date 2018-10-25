package service

import "github.com/noppawitt/admintools/repository"

// ParameterService provides function service
type ParameterService interface {
	Repo() repository.ParameterRepository
}

type parameterService struct {
	Repository repository.ParameterRepository
}

// NewParameterService returns parameter service
func NewParameterService(r repository.ParameterRepository) ParameterService {
	return &parameterService{r}
}

func (s *parameterService) Repo() repository.ParameterRepository {
	return s.Repository
}
