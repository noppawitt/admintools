package service

import (
	"fmt"

	"github.com/noppawitt/admintools/model"
	"github.com/noppawitt/admintools/repository"
	"github.com/noppawitt/admintools/util"
)

// ApplicationService provides application service
type ApplicationService interface {
	Repo() repository.ApplicationRepository
	Create(application *model.Application) error
	getConnectionString(application *model.Application) (string, error)
	GetStoredProcedure(id int) ([]model.ExternalStoredProcedure, error)
	GetView(id int) ([]model.ExternalView, error)
	GetParameter(id int, name string) ([]model.ExternalParameter, error)
}

type applicationService struct {
	Repository         repository.ApplicationRepository
	ExternalRepository repository.ExternalRepository
	EncryptionKey      string
}

// NewApplicationService returns application service
func NewApplicationService(r repository.ApplicationRepository, e repository.ExternalRepository, encryptionKey string) ApplicationService {
	return &applicationService{r, e, encryptionKey}
}

func (s *applicationService) Repo() repository.ApplicationRepository {
	return s.Repository
}

func (s *applicationService) Create(application *model.Application) error {
	var err error
	encryptedPassword, err := util.Encrypt(application.Password, s.EncryptionKey)
	if err != nil {
		return err
	}
	application.Password = encryptedPassword
	err = s.Repository.Create(application)
	return err
}

func (s *applicationService) getConnectionString(application *model.Application) (string, error) {
	decryptedPassword, err := util.Decrypt(application.Password, s.EncryptionKey)
	if err != nil {
		return "", err
	}
	cs := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable", application.Username, decryptedPassword, application.Host, application.Port, application.DBName)
	return cs, err
}

func (s *applicationService) GetStoredProcedure(id int) ([]model.ExternalStoredProcedure, error) {
	application, err := s.Repository.FindOne(id)
	cs, err := s.getConnectionString(application)
	storedProcedures, err := s.ExternalRepository.FindStoredProcedure(cs)
	return storedProcedures, err
}

func (s *applicationService) GetView(id int) ([]model.ExternalView, error) {
	application, err := s.Repository.FindOne(id)
	cs, err := s.getConnectionString(application)
	views, err := s.ExternalRepository.FindView(cs)
	return views, err
}

func (s *applicationService) GetParameter(id int, name string) ([]model.ExternalParameter, error) {
	application, err := s.Repository.FindOne(id)
	cs, err := s.getConnectionString(application)
	parameters, err := s.ExternalRepository.FindParameter(cs, name)
	return parameters, err
}
