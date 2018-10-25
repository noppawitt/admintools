package model

import "time"

// Application is a application model
type Application struct {
	ID        int        `gorm:"primary_key" json:"id" valid:"int,optional"`
	Name      string     `json:"name" valid:"required"`
	Host      string     `json:"host" valid:"ipv4,required"`
	Port      int        `json:"port" valid:"port,required"`
	Username  string     `json:"username" valid:"username,required"`
	Password  string     `json:"password" valid:"required"`
	DBName    string     `json:"dbName" valid:"required"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
	Functions []Function `gorm:"foreignkey:ApplicationID" json:"functions"`
}
