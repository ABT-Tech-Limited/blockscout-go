package blockscout

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetStats(t *testing.T) {
	stats, err := api.GetStats(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, stats)
	assert.NotEmpty(t, stats.TotalTransactions)
	assert.NotEmpty(t, stats.TotalBlocks)
}

func TestClient_GetTransactionsChart(t *testing.T) {
	chart, err := api.GetTransactionsChart(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, chart)
}

func TestClient_GetMarketChart(t *testing.T) {
	chart, err := api.GetMarketChart(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, chart)
}
