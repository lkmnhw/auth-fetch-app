package handlers

import (
	"context"
	"fetch-app/models"

	cache "github.com/patrickmn/go-cache"
)

type Handler struct {
	Context context.Context
	Cache   *cache.Cache
}

type Response struct {
	Commodities []*models.Commodity `json:"commodities,omitempty"`
	Aggregates  []*models.Aggregate `json:"aggregates,omitempty"`
	Message     string              `json:"message"`
}
