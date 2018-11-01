package service

import (
	"testing"

	"github.com/noppawitt/admintools/model"
	"github.com/noppawitt/admintools/repository"
	"github.com/noppawitt/admintools/util"

	"github.com/stretchr/testify/assert"

	"github.com/noppawitt/admintools/repository/mocks"
)

const key = "11111111111111111111111111111111"

func TestRepo(t *testing.T) {
	r := &mocks.ApplicationRepository{}
	e := &mocks.ExternalRepository{}
	applicationService := NewApplicationService(r, e, key)
	repo := applicationService.Repo()
	assert.Implements(t, (*repository.ApplicationRepository)(nil), repo)
}

// func TestCreate(t *testing.T) {
// 	r := &mocks.ApplicationRepository{}
// 	r.On("Create", &model.Application, func(application *model.Application) error {
// 		application.ID = 1
// 		application.CreatedAt = time.Now()
// 		application.UpdatedAt = time.Now()
// 		return nil
// 	})
// 	e := mock.ExternalRepository{}
// 	applicationService := NewApplicationService(r, e, key)
// }

func TestGetConnectionString(t *testing.T) {
	r := &mocks.ApplicationRepository{}
	e := &mocks.ExternalRepository{}
	applicationService := NewApplicationService(r, e, key)

	password, _ := util.Encrypt("password", key)

	application := &model.Application{
		Name:     "app",
		Host:     "127.0.0.1",
		Port:     1433,
		Username: "admin",
		Password: password,
		DBName:   "db",
	}

	cs, err := applicationService.getConnectionString(application)
	assert.Equal(t, nil, err)
	assert.Equal(t, "sqlserver://admin:password@127.0.0.1:1433?database=db&encrypt=disable", cs)

	applicationWithOutPassword := &model.Application{
		Name:     "app",
		Host:     "127.0.0.1",
		Port:     1433,
		Username: "admin",
		Password: "",
		DBName:   "db",
	}
	cs, err = applicationService.getConnectionString(applicationWithOutPassword)
	assert.Error(t, err)
}
