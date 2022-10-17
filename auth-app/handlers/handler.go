package handlers

import (
	"auth-app/databases"
	"auth-app/models"
	"context"
)

type Handler struct {
	Context  context.Context
	Database *databases.Database
}

type Response struct {
	User    *models.User `json:"user,omitempty"`
	Message string       `json:"message"`
}
