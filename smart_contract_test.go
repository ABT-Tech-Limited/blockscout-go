package blockscout

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_ListSmartContracts(t *testing.T) {
	resp, err := api.ListSmartContracts(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestClient_GetSmartContract(t *testing.T) {
	list, err := api.ListSmartContracts(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, list)
	if len(list.Items) == 0 {
		t.Skip("no smart contracts found, skipping")
	}

	sc, err := api.GetSmartContract(context.Background(), list.Items[0].Address.Hash)
	assert.NoError(t, err)
	assert.NotNil(t, sc)
}

func TestClient_GetSmartContractCounters(t *testing.T) {
	counters, err := api.GetSmartContractCounters(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, counters)
}
