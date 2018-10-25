package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/noppawitt/admintools/model"
)

// ApplicationRepository provides access a application store
type ApplicationRepository interface {
	Create(application *model.Application) error
	Find() ([]model.Application, error)
	FindOne(id int) (*model.Application, error)
	Save(application *model.Application) error
	Remove(id int) error
	Paginate(name string, dbName string, sortBy string, direction string, limit int, offset int) ([]model.Application, error)
	Count(name string, dbName string) (int, error)
}

type applicationRepository struct {
	DB *gorm.DB
}

// NewApplicationRepository returns application repository
func NewApplicationRepository(db *gorm.DB) ApplicationRepository {
	return &applicationRepository{db}
}

func (r *applicationRepository) Create(application *model.Application) error {
	err := r.DB.Create(application).Error
	return err
}

func (r *applicationRepository) Find() ([]model.Application, error) {
	applications := []model.Application{}
	err := r.DB.Find(&applications).Error
	return applications, err
}

func (r *applicationRepository) FindOne(id int) (*model.Application, error) {
	application := &model.Application{}
	err := r.DB.Find(application, id).Error
	return application, err
}

func (r *applicationRepository) Save(application *model.Application) error {
	err := r.DB.Save(application).Error
	return err
}

func (r *applicationRepository) Remove(id int) error {
	err := r.DB.Where("id = ?", id).Delete(&model.Application{}).Error
	return err
}

func (r *applicationRepository) Paginate(name string, dbName string, sortBy string, direction string, limit int, offset int) ([]model.Application, error) {
	applications := []model.Application{}
	if limit == 0 {
		limit = 10
	}
	query := r.DB
	if name != "" {
		query = query.Where("name like ?", "%"+name+"%")
	}
	if dbName != "" {
		query = query.Where("db_name like ?", "%"+dbName+"%")
	}
	query = query.Limit(limit).Offset(offset).Order(fmt.Sprintf("%s %s", sortBy, direction))
	err := query.Find(&applications).Error
	return applications, err
}

func (r *applicationRepository) Count(name string, dbName string) (int, error) {
	count := 0
	query := r.DB.Model(&model.Application{})
	if name != "" {
		query = query.Where("name like ?", "%"+name+"%")
	}
	if dbName != "" {
		query = query.Where("db_name like ?", "%"+dbName+"%")
	}
	err := query.Count(&count).Error
	return count, err
}
