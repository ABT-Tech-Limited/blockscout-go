package blockscout

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Search(t *testing.T) {
	resp, err := api.Search(context.Background(), &SearchParams{Q: "0x"})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestClient_CheckRedirect(t *testing.T) {
	// Use a block number to test redirect
	resp, err := api.CheckRedirect(context.Background(), "1")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
