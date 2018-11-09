package service

import (
	"testing"

	"github.com/noppawitt/admintools/repository"

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
