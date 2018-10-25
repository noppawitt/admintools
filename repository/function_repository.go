package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/noppawitt/admintools/model"
)

// FunctionRepository provides access a function store
type FunctionRepository interface {
	Create(function *model.Function) error
	FindOne(id int) (*model.Function, error)
	FindOneIncludeApplication(id int) (*model.Function, error)
	Find() ([]model.Function, error)
	Save(function *model.Function) error
	Remove(id int) error
	Paginate(name string, appName string, sortBy string, direction string, limit int, offset int) ([]model.Function, error)
	Count(name string, appName string) (int, error)
}

type functionRepository struct {
	DB *gorm.DB
}

// NewFunctionRepository returns function repository
func NewFunctionRepository(db *gorm.DB) FunctionRepository {
	return &functionRepository{db}
}

func (r functionRepository) Create(function *model.Function) error {
	err := r.DB.Create(function).Error
	return err
}

func (r functionRepository) Find() ([]model.Function, error) {
	functions := []model.Function{}
	err := r.DB.Find(&functions).Error
	return functions, err
}

func (r functionRepository) FindOne(id int) (*model.Function, error) {
	function := &model.Function{}
	err := r.DB.Find(function).Error
	return function, err
}

func (r functionRepository) FindOneIncludeApplication(id int) (*model.Function, error) {
	function := &model.Function{}
	err := r.DB.Preload("Application").Find(&function, id).Error
	return function, err
}

func (r functionRepository) Save(function *model.Function) error {
	err := r.DB.Save(function).Error
	return err
}

func (r functionRepository) Remove(id int) error {
	err := r.DB.Where("id = ?", id).Delete(&model.Function{}).Error
	return err
}

func (r *functionRepository) Paginate(name string, appName string, sortBy string, direction string, limit int, offset int) ([]model.Function, error) {
	functions := []model.Function{}
	if limit == 0 {
		limit = 10
	}
	if sortBy == "name" {
		sortBy = "functions.name"
	} else if sortBy == "app_name" {
		sortBy = "applications.name"
	}
	err := r.DB.Preload("Application").
		Joins("join applications on functions.application_id = applications.id").
		Where("functions.name like ?", "%"+name+"%").
		Where("applications.name like ?", "%"+appName+"%").
		Limit(limit).
		Offset(offset).
		Order(sortBy + " " + direction).
		Find(&functions).Error
	return functions, err
}

func (r *functionRepository) Count(name string, appName string) (int, error) {
	count := 0
	err := r.DB.Table("functions").
		Joins("join applications on functions.application_id = applications.id").
		Where("functions.name like ?", "%"+name+"%").
		Where("applications.name like ?", "%"+appName+"%").
		Count(&count).Error
	return count, err
}
