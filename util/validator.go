package util

import (
	valid "github.com/asaskevich/govalidator"
)

// RegisterCustomValidator registers custom validator
func RegisterCustomValidator() {
	valid.TagMap["username"] = valid.Validator(func(str string) bool {
		return valid.Matches(str, "^[0-9A-Za-z_-]+$")
	})
}
