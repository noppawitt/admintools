package model

import "time"

// Parameter is a parameter model
type Parameter struct {
	ID          int        `gorm:"primary_key" json:"id"`
	Name        string     `json:"name"`
	DisplayName string     `json:"displayName"`
	Type        string     `json:"type"`
	Length      int        `json:"length"`
	FunctionID  uint       `json:"functionID"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"-"`
	Function    Function   `gorm:"foreignkey:FunctionID" json:"function"`
}
