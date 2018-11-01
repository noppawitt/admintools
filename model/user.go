package model

// User is an user model
type User struct {
	ID              string `gorm:"primary_key" json:"id"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	RefreshToken    string `json:"refreshToken"`
	SSORefreshToken string `json:"ssoRefreshToken"`
}
