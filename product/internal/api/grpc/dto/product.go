package dto

import "github.com/google/uuid"

type SearchProductsDTO struct {
	UserID     uuid.UUID
	CategoryID int32
	Moderated  bool
	ProductIDs []uuid.UUID
}
