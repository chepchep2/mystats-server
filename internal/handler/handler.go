package handler

import (
	"mystats-server/internal/ent"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	client    *ent.Client
	validator *validator.Validate
}

func NewHandler(client *ent.Client) *Handler {
	return &Handler{
		client:    client,
		validator: validator.New(),
	}
}
