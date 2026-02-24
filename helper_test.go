package blockscout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildPath(t *testing.T) {
	result := buildPath("/api/v2/addresses/{address_hash}/transactions", map[string]string{
		"address_hash": "0x1234567890abcdef",
	})
	assert.Equal(t, "/api/v2/addresses/0x1234567890abcdef/transactions", result)
}

func TestBuildPath_MultipleParams(t *testing.T) {
	result := buildPath("/api/v2/tokens/{address_hash}/instances/{token_id}", map[string]string{
		"address_hash": "0xabc",
		"token_id":     "42",
	})
	assert.Equal(t, "/api/v2/tokens/0xabc/instances/42", result)
}

func TestStructToQueryParams(t *testing.T) {
	type Params struct {
		Name  string `query:"name"`
		Value string `query:"value,omitempty"`
		Count *int   `query:"count,omitempty"`
	}

	count := 10
	params := structToQueryParams(&Params{
		Name:  "test",
		Value: "hello",
		Count: &count,
	})
	assert.Equal(t, "test", params["name"])
	assert.Equal(t, "hello", params["value"])
	assert.Equal(t, "10", params["count"])
}

func TestStructToQueryParams_Omitempty(t *testing.T) {
	type Params struct {
		Name  string `query:"name"`
		Value string `query:"value,omitempty"`
		Count *int   `query:"count,omitempty"`
	}

	params := structToQueryParams(&Params{
		Name: "test",
	})
	assert.Equal(t, "test", params["name"])
	_, hasValue := params["value"]
	assert.False(t, hasValue)
	_, hasCount := params["count"]
	assert.False(t, hasCount)
}

func TestStructToQueryParams_Nil(t *testing.T) {
	params := structToQueryParams(nil)
	assert.Nil(t, params)
}
