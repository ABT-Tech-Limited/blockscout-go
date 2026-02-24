package blockscout

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_ListTransactions(t *testing.T) {
	resp, err := api.ListTransactions(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Greater(t, len(resp.Items), 0)
}

func TestClient_GetTransaction(t *testing.T) {
	list, err := api.ListTransactions(context.Background(), nil)
	assert.NoError(t, err)
	assert.Greater(t, len(list.Items), 0)

	tx, err := api.GetTransaction(context.Background(), list.Items[0].Hash)
	assert.NoError(t, err)
	assert.NotNil(t, tx)
	assert.Equal(t, list.Items[0].Hash, tx.Hash)
}

func TestClient_GetTransactionStats(t *testing.T) {
	stats, err := api.GetTransactionStats(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, stats)
}

func TestClient_ListTransactionLogs(t *testing.T) {
	list, err := api.ListTransactions(context.Background(), nil)
	assert.NoError(t, err)
	assert.Greater(t, len(list.Items), 0)

	resp, err := api.ListTransactionLogs(context.Background(), list.Items[0].Hash, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
