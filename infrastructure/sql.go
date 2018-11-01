package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/noppawitt/admintools/model"
)

// Connect returns a new db instance
func Connect(dialect string, cs string) (*gorm.DB, error) {
	db, err := gorm.Open(dialect, cs)
	return db, err
}

// AutoMigrate migrates all models to database
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Application{})
	db.AutoMigrate(&model.Function{})
	db.AutoMigrate(&model.Parameter{})
}
