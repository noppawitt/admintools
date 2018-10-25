package repository

import (
	"github.com/noppawitt/admintools/infrastructure"
	"github.com/noppawitt/admintools/model"
)

// ExternalRepository provides access a external store
type ExternalRepository interface {
	FindStoredProcedure(cs string) ([]model.ExternalStoredProcedure, error)
	FindView(cs string) ([]model.ExternalView, error)
	FindParameter(cs string, name string) ([]model.ExternalParameter, error)
}

type externalRepository struct {
}

// NewExternalRepository returns external repository
func NewExternalRepository() ExternalRepository {
	return &externalRepository{}
}

func (r *externalRepository) FindStoredProcedure(cs string) ([]model.ExternalStoredProcedure, error) {
	db, err := infrastructure.Connect("mssql", cs)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	storedProcedures := []model.ExternalStoredProcedure{}
	err = db.Raw("select name from sys.procedures").Scan(&storedProcedures).Error
	return storedProcedures, err
}

func (r *externalRepository) FindView(cs string) ([]model.ExternalView, error) {
	db, err := infrastructure.Connect("mssql", cs)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	views := []model.ExternalView{}
	err = db.Raw("select name from sys.views").Scan(&views).Error
	return views, err
}

func (r *externalRepository) FindParameter(cs string, name string) ([]model.ExternalParameter, error) {
	db, err := infrastructure.Connect("mssql", cs)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	parameters := []model.ExternalParameter{}
	err = db.Raw(`
		select
			name,
			type_name(user_type_id) as type,
			max_length as length
		from sys.parameters
		where object_id = object_id(?)`,
		name,
	).Scan(&parameters).Error
	return parameters, err
}
