package blockscout

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_ListAddresses(t *testing.T) {
	resp, err := api.ListAddresses(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Greater(t, len(resp.Items), 0)
}

func TestClient_GetAddress(t *testing.T) {
	// First get an address from the list
	list, err := api.ListAddresses(context.Background(), nil)
	assert.NoError(t, err)
	assert.Greater(t, len(list.Items), 0)

	addr, err := api.GetAddress(context.Background(), list.Items[0].Hash)
	assert.NoError(t, err)
	assert.NotNil(t, addr)
	assert.Equal(t, list.Items[0].Hash, addr.Hash)
}

func TestClient_ListAddressTransactions(t *testing.T) {
	list, err := api.ListAddresses(context.Background(), nil)
	assert.NoError(t, err)
	assert.Greater(t, len(list.Items), 0)

	resp, err := api.ListAddressTransactions(context.Background(), list.Items[0].Hash, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestClient_GetAddressCounters(t *testing.T) {
	list, err := api.ListAddresses(context.Background(), nil)
	assert.NoError(t, err)
	assert.Greater(t, len(list.Items), 0)

	counters, err := api.GetAddressCounters(context.Background(), list.Items[0].Hash)
	assert.NoError(t, err)
	assert.NotNil(t, counters)
}

func TestClient_GetAddressTabsCounters(t *testing.T) {
	list, err := api.ListAddresses(context.Background(), nil)
	assert.NoError(t, err)
	assert.Greater(t, len(list.Items), 0)

	counters, err := api.GetAddressTabsCounters(context.Background(), list.Items[0].Hash)
	assert.NoError(t, err)
	assert.NotNil(t, counters)
}
