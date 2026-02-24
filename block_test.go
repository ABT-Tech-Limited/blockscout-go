package blockscout

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_ListBlocks(t *testing.T) {
	resp, err := api.ListBlocks(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Greater(t, len(resp.Items), 0)
}

func TestClient_GetBlock(t *testing.T) {
	list, err := api.ListBlocks(context.Background(), nil)
	assert.NoError(t, err)
	assert.Greater(t, len(list.Items), 0)

	block, err := api.GetBlock(context.Background(), list.Items[0].Hash)
	assert.NoError(t, err)
	assert.NotNil(t, block)
	assert.Equal(t, list.Items[0].Hash, block.Hash)
}

func TestClient_ListBlockTransactions(t *testing.T) {
	list, err := api.ListBlocks(context.Background(), nil)
	assert.NoError(t, err)
	assert.Greater(t, len(list.Items), 0)

	resp, err := api.ListBlockTransactions(context.Background(), list.Items[0].Hash, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
