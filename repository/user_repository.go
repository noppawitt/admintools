package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/noppawitt/admintools/model"
)

// UserRepository provides access a user store
type UserRepository interface {
	Create(user *model.User) error
	FindOne(id string) (*model.User, error)
	Save(user *model.User) error
}

type userRepository struct {
	DB *gorm.DB
}

// NewUserRepository returns new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	err := r.DB.Create(user).Error
	return err
}

func (r *userRepository) FindOne(id string) (*model.User, error) {
	user := &model.User{}
	err := r.DB.First(user, "id = ?", id).Error
	return user, err
}

func (r *userRepository) Save(user *model.User) error {
	err := r.DB.Save(user).Error
	return err
}
