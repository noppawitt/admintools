package model

import "time"

// Function is a function model
type Function struct {
	ID                  int         `gorm:"primary_key" json:"id"`
	Name                string      `json:"name"`
	Remarks             string      `json:"remarks"`
	AllowMultiple       bool        `json:"allowMultiple"`
	StoredProcedureName string      `json:"storedProcedureName"`
	ViewName            string      `json:"viewName"`
	ApplicationID       int         `json:"applicationID"`
	CreatedAt           time.Time   `json:"createdAt"`
	UpdatedAt           time.Time   `json:"updatedAt"`
	DeletedAt           *time.Time  `json:"-"`
	Application         Application `gorm:"foreignkey:ApplicationID" json:"application"`
	Parameters          []Parameter `gorm:"foreignkey:FunctionID" json:"parameters"`
}
