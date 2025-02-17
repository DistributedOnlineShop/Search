package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"Search/util"
)

func CreateRandomSearchIndex(t *testing.T) SearchIndex {
	var jsondata *gofakeit.JSONOptions
	json, err := gofakeit.JSON(jsondata)
	require.NoError(t, err)

	data := CreateSearchIndexParams{
		ProductID:         util.GenerateProductID(),
		ProductName:       gofakeit.ProductName(),
		ProductDesc:       gofakeit.ProductDescription(),
		ProductAttributes: json,
	}

	si, err := testStore.CreateSearchIndex(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, si)
	require.Equal(t, si.ProductID, data.ProductID)
	require.Equal(t, si.ProductName, data.ProductName)
	require.Equal(t, si.ProductDesc, data.ProductDesc)
	require.Equal(t, si.ProductAttributes, data.ProductAttributes)
	require.NotEmpty(t, si.IndexedAt)

	return si
}

func TestCreateSearchIndex(t *testing.T) {
	CreateRandomSearchIndex(t)
}

func TestGetSearchIndexByProductId(t *testing.T) {
	si := CreateRandomSearchIndex(t)

	items, err := testStore.GetSearchIndexByProductId(context.Background(), si.ProductID)
	require.NoError(t, err)
	require.NotEmpty(t, items)
}
