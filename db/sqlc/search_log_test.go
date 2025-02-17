package db

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"

	"Search/util"
)

func CreateRandomSearchLog(t *testing.T) SearchLog {
	var jsondata *gofakeit.JSONOptions
	json, err := gofakeit.JSON(jsondata)
	require.NoError(t, err)

	data := CreateSearchLogParams{
		SearchID:      util.CreateUUID(),
		UserID:        util.CreateUUID(),
		SearchQuery:   gofakeit.Word(),
		SearchFilters: json,
		ResultsCount:  util.GenerateInt32(),
	}

	sl, err := testStore.CreateSearchLog(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, sl)
	require.Equal(t, sl.SearchID, data.SearchID)
	require.Equal(t, sl.UserID, data.UserID)
	require.Equal(t, sl.SearchQuery, data.SearchQuery)
	require.Equal(t, sl.SearchFilters, data.SearchFilters)
	require.Equal(t, sl.ResultsCount, data.ResultsCount)
	require.NotEmpty(t, sl.SearchedAt)

	return sl
}

func TestCreateSearchLog(t *testing.T) {
	CreateRandomSearchLog(t)
}

func TestGetSearchLogByTime(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomSearchLog(t)
	}

	time := GetSearchLogByTimeParams{
		SearchedAt:   pgtype.Timestamp{Time: time.Now().Add(-20000 * time.Hour), Valid: true},
		SearchedAt_2: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	items, err := testStore.GetSearchLogByTime(context.Background(), time)
	require.NoError(t, err)
	require.NotEmpty(t, items)
}

func TestGetSearchLogByUserId(t *testing.T) {
	sl := CreateRandomSearchLog(t)

	items, err := testStore.GetSearchLogByUserId(context.Background(), sl.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, items)
}
func TestGetSearchLogList(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomSearchLog(t)
	}

	items, err := testStore.GetSearchLogList(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.GreaterOrEqual(t, len(items), 10)
}
