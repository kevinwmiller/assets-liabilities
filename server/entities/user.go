package entities

import "github.com/microcosm-cc/bluemonday"

// User represents a user of the platform. Could store additional user information such as email in the future
type User struct {
	Entity
	Username string
	Password string
	FullName string
}

// SanitizeUser sanitizes potential HTML elements in the model's string fields
func SanitizeUser(data *User) {
	strict := bluemonday.StrictPolicy()
	data.Username = strict.Sanitize(data.Username)
	data.FullName = strict.Sanitize(data.FullName)
}
