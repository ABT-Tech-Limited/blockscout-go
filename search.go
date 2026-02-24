package blockscout

import "context"

// ── REQUEST TYPES ──

// SearchParams represents query parameters for the search endpoint.
type SearchParams struct {
	Q          string `query:"q,omitempty"`
	ItemsCount *int   `query:"items_count,omitempty"`
}

// ── RESPONSE TYPES ──

// SearchResult represents a search result item.
type SearchResult struct {
	AddressHash              string  `json:"address_hash"`
	AddressURL               string  `json:"address_url"`
	Certified                bool    `json:"certified"`
	CirculatingMarketCap     *string `json:"circulating_market_cap"`
	ExchangeRate             *string `json:"exchange_rate"`
	IconURL                  *string `json:"icon_url"`
	IsSmartContractVerified  bool    `json:"is_smart_contract_verified"`
	IsVerifiedViaAdminPanel  bool    `json:"is_verified_via_admin_panel"`
	Name                     string  `json:"name"`
	Priority                 int     `json:"priority"`
	Symbol                   string  `json:"symbol"`
	TokenType                string  `json:"token_type"`
	TokenURL                 string  `json:"token_url"`
	TotalSupply              string  `json:"total_supply"`
	Type                     string  `json:"type"` // "token" | "address" | "contract" | "block" | "transaction"
	// Block-specific fields
	BlockHash                string  `json:"block_hash,omitempty"`
	BlockNumber              *int    `json:"block_number,omitempty"`
	BlockType                string  `json:"block_type,omitempty"`
	Timestamp                *string `json:"timestamp,omitempty"`
	// Transaction-specific fields
	TransactionHash          string  `json:"tx_hash,omitempty"`
	URL                      string  `json:"url,omitempty"`
}

// SearchRedirect represents a search redirect response.
type SearchRedirect struct {
	Parameter *string `json:"parameter"`
	Redirect  *bool   `json:"redirect"`
	Type      *string `json:"type"` // "address" | "block" | "transaction" | "user_operation" | "blob"
}

// ── API METHODS ──

func (c *client) Search(ctx context.Context, params *SearchParams) (resp *PaginatedResponse[SearchResult], err error) {
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get("/api/v2/search")
	return
}

func (c *client) CheckRedirect(ctx context.Context, query string) (resp *SearchRedirect, err error) {
	_, err = c.resty.R().
		SetContext(ctx).
		SetQueryParam("q", query).
		SetResult(&resp).
		Get("/api/v2/search/check-redirect")
	return
}
