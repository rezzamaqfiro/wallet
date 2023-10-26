package order

import (
	repo "github.com/rezzamaqfiro/wallet/repo/generated"
)

type Handler struct {
	db *repo.Queries
}

func NewHandler(db *repo.Queries) *Handler {
	return &Handler{db}
}
