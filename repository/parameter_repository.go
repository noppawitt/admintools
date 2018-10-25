package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/noppawitt/admintools/model"
)

// ParameterRepository provides access a parameter store
type ParameterRepository interface {
	Create(parameter *model.Parameter) error
	Find() ([]model.Parameter, error)
	FindOne(id int) (*model.Parameter, error)
	Save(parameter *model.Parameter) error
	Remove(id int) error
	FindByFunctionID(functionID int) ([]model.Parameter, error)
}

type parameterRepository struct {
	DB *gorm.DB
}

// NewParameterRepository returns parameter repsitory
func NewParameterRepository(db *gorm.DB) ParameterRepository {
	return &parameterRepository{db}
}

func (r parameterRepository) Create(parameter *model.Parameter) error {
	err := r.DB.Create(parameter).Error
	return err
}

func (r parameterRepository) Find() ([]model.Parameter, error) {
	parameters := []model.Parameter{}
	err := r.DB.Find(&parameters).Error
	return parameters, err
}

func (r parameterRepository) FindOne(id int) (*model.Parameter, error) {
	parameter := &model.Parameter{}
	err := r.DB.Find(parameter, id).Error
	return parameter, err
}

func (r parameterRepository) Save(parameter *model.Parameter) error {
	err := r.DB.Save(parameter).Error
	return err
}

func (r parameterRepository) Remove(id int) error {
	err := r.DB.Where("id = ?", id).Delete(&model.Parameter{}).Error
	return err
}

func (r parameterRepository) FindByFunctionID(functionID int) ([]model.Parameter, error) {
	parameters := []model.Parameter{}
	err := r.DB.Where("function_id = ?", functionID).Find(&parameters).Error
	return parameters, err
}
