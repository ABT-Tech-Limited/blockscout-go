package blockscout

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_ListTokens(t *testing.T) {
	resp, err := api.ListTokens(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestClient_GetToken(t *testing.T) {
	list, err := api.ListTokens(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, list)
	if len(list.Items) == 0 {
		t.Skip("no tokens found, skipping")
	}

	token, err := api.GetToken(context.Background(), list.Items[0].AddressHash)
	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.Equal(t, list.Items[0].AddressHash, token.AddressHash)
}

func TestClient_GetTokenCounters(t *testing.T) {
	list, err := api.ListTokens(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, list)
	if len(list.Items) == 0 {
		t.Skip("no tokens found, skipping")
	}

	counters, err := api.GetTokenCounters(context.Background(), list.Items[0].AddressHash)
	assert.NoError(t, err)
	assert.NotNil(t, counters)
}

func TestClient_ListTokenHolders(t *testing.T) {
	list, err := api.ListTokens(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, list)
	if len(list.Items) == 0 {
		t.Skip("no tokens found, skipping")
	}

	resp, err := api.ListTokenHolders(context.Background(), list.Items[0].AddressHash, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
