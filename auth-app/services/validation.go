package services

import (
	"auth-app/models"
)

func userValidation(user models.User) string {
	if user.Phone == "" {
		return `phone required`
	}

	if user.Name == "" {
		return `name required`
	}

	if user.Role == "" {
		return `role required`
	}

	return ""
}
