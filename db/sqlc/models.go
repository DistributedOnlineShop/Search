// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type SearchIndex struct {
	ProductID         string           `json:"product_id"`
	ProductName       string           `json:"product_name"`
	ProductDesc       pgtype.Text      `json:"product_desc"`
	ProductAttributes []byte           `json:"product_attributes"`
	IndexedAt         pgtype.TIMESTAMP(0) `json:"indexed_at"`
}

type SearchLog struct {
	SearchID      uuid.UUID        `json:"search_id"`
	UserID        pgtype.UUID      `json:"user_id"`
	SearchQuery   string           `json:"search_query"`
	SearchFilters []byte           `json:"search_filters"`
	ResultsCount  int32            `json:"results_count"`
	SearchedAt    pgtype.TIMESTAMP(0) `json:"searched_at"`
}
