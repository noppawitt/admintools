package service

import "github.com/noppawitt/admintools/repository"

// FunctionService provides function service
type FunctionService interface {
	Repo() repository.FunctionRepository
}

type functionService struct {
	Repository repository.FunctionRepository
}

// NewFunctionService returns function service
func NewFunctionService(r repository.FunctionRepository) FunctionService {
	return &functionService{r}
}

func (s *functionService) Repo() repository.FunctionRepository {
	return s.Repository
}
