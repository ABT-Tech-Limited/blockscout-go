package blockscout

import "context"

// ── REQUEST TYPES ──

// ListTokensParams represents query parameters for listing tokens.
type ListTokensParams struct {
	Type                 string `query:"type,omitempty"` // ERC-20, ERC-721, ERC-1155, ERC-404
	Q                    string `query:"q,omitempty"`    // search by name/symbol
	Sort                 string `query:"sort,omitempty"`
	Order                string `query:"order,omitempty"` // "asc" | "desc"
	ContractAddressHash  string `query:"contract_address_hash,omitempty"`
	HoldersCount         *int   `query:"holders_count,omitempty"`
	FiatValue            string `query:"fiat_value,omitempty"`
	MarketCap            string `query:"market_cap,omitempty"`
	IsNameNull           *bool  `query:"is_name_null,omitempty"`
	Name                 string `query:"name,omitempty"`
	ItemsCount           *int   `query:"items_count,omitempty"`
}

// ── RESPONSE TYPES ──

// TokenHolder represents a token holder entry.
type TokenHolder struct {
	Address  Address `json:"address"`
	Value    string  `json:"value"`
	TokenID  *string `json:"token_id"`
}

// TokenCounters represents token statistics.
type TokenCounters struct {
	TokenHoldersCount string `json:"token_holders_count"`
	TransfersCount    string `json:"transfers_count"`
}

// ── API METHODS ──

func (c *client) ListTokens(ctx context.Context, params *ListTokensParams) (resp *PaginatedResponse[Token], err error) {
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get("/api/v2/tokens")
	return
}

func (c *client) GetToken(ctx context.Context, addressHash string) (resp *Token, err error) {
	path := buildPath("/api/v2/tokens/{address_hash}", map[string]string{
		"address_hash": addressHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) ListTokenTransfers(ctx context.Context, addressHash string, params *PaginationParams) (resp *PaginatedResponse[TokenTransfer], err error) {
	path := buildPath("/api/v2/tokens/{address_hash}/transfers", map[string]string{
		"address_hash": addressHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) ListTokenHolders(ctx context.Context, addressHash string, params *PaginationParams) (resp *PaginatedResponse[TokenHolder], err error) {
	path := buildPath("/api/v2/tokens/{address_hash}/holders", map[string]string{
		"address_hash": addressHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) GetTokenCounters(ctx context.Context, addressHash string) (resp *TokenCounters, err error) {
	path := buildPath("/api/v2/tokens/{address_hash}/counters", map[string]string{
		"address_hash": addressHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) ListTokenInstances(ctx context.Context, addressHash string, params *PaginationParams) (resp *PaginatedResponse[TokenInstance], err error) {
	path := buildPath("/api/v2/tokens/{address_hash}/instances", map[string]string{
		"address_hash": addressHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) GetTokenInstance(ctx context.Context, addressHash string, tokenID string) (resp *TokenInstance, err error) {
	path := buildPath("/api/v2/tokens/{address_hash}/instances/{token_id}", map[string]string{
		"address_hash": addressHash,
		"token_id":     tokenID,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) ListTokenInstanceTransfers(ctx context.Context, addressHash string, tokenID string, params *PaginationParams) (resp *PaginatedResponse[TokenTransfer], err error) {
	path := buildPath("/api/v2/tokens/{address_hash}/instances/{token_id}/transfers", map[string]string{
		"address_hash": addressHash,
		"token_id":     tokenID,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) ListTokenInstanceHolders(ctx context.Context, addressHash string, tokenID string, params *PaginationParams) (resp *PaginatedResponse[TokenHolder], err error) {
	path := buildPath("/api/v2/tokens/{address_hash}/instances/{token_id}/holders", map[string]string{
		"address_hash": addressHash,
		"token_id":     tokenID,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}
