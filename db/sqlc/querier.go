// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateSearchIndex(ctx context.Context, arg CreateSearchIndexParams) (SearchIndex, error)
	CreateSearchLog(ctx context.Context, arg CreateSearchLogParams) (SearchLog, error)
	GetSearchIndexByProductId(ctx context.Context, productID string) ([]SearchIndex, error)
	GetSearchLogByTime(ctx context.Context, arg GetSearchLogByTimeParams) ([]SearchLog, error)
	GetSearchLogByUserId(ctx context.Context, userID uuid.UUID) ([]SearchLog, error)
	GetSearchLogList(ctx context.Context) ([]SearchLog, error)
}

var _ Querier = (*Queries)(nil)
