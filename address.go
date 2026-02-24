package blockscout

import "context"

// ── REQUEST TYPES ──

// ListAddressesParams represents query parameters for listing addresses.
type ListAddressesParams struct {
	Sort              string `query:"sort,omitempty"`               // "balance" | "transactions_count"
	Order             string `query:"order,omitempty"`              // "asc" | "desc"
	FetchedCoinBalance *string `query:"fetched_coin_balance,omitempty"`
	Hash              string `query:"hash,omitempty"`
	ItemsCount        *int   `query:"items_count,omitempty"`
	TransactionsCount *int   `query:"transactions_count,omitempty"`
}

// ListAddressTransactionsParams represents query parameters for listing address transactions.
type ListAddressTransactionsParams struct {
	Filter      string  `query:"filter,omitempty"`       // "to" | "from"
	Sort        string  `query:"sort,omitempty"`          // "block_number" | "value" | "fee"
	Order       string  `query:"order,omitempty"`         // "asc" | "desc"
	BlockNumber *int    `query:"block_number,omitempty"`
	Index       *int    `query:"index,omitempty"`
	InsertedAt  string  `query:"inserted_at,omitempty"`
	Hash        string  `query:"hash,omitempty"`
	Value       string  `query:"value,omitempty"`
	Fee         string  `query:"fee,omitempty"`
	ItemsCount  *int    `query:"items_count,omitempty"`
}

// ListAddressTokenTransfersParams represents query parameters for listing address token transfers.
type ListAddressTokenTransfersParams struct {
	Filter               string `query:"filter,omitempty"`                 // "to" | "from"
	Type                 string `query:"type,omitempty"`                   // comma-separated: ERC-20, ERC-721, ERC-1155, ERC-404
	Token                string `query:"token,omitempty"`                  // token contract address
	BlockNumber          *int   `query:"block_number,omitempty"`
	Index                *int   `query:"index,omitempty"`
	ItemsCount           *int   `query:"items_count,omitempty"`
	BatchLogIndex        *int   `query:"batch_log_index,omitempty"`
	BatchBlockHash       string `query:"batch_block_hash,omitempty"`
	BatchTransactionHash string `query:"batch_transaction_hash,omitempty"`
	IndexInBatch         *int   `query:"index_in_batch,omitempty"`
}

// ListAddressInternalTransactionsParams represents query parameters for listing address internal transactions.
type ListAddressInternalTransactionsParams struct {
	Filter           string `query:"filter,omitempty"` // "to" | "from"
	BlockNumber      *int   `query:"block_number,omitempty"`
	Index            *int   `query:"index,omitempty"`
	TransactionIndex *int   `query:"transaction_index,omitempty"`
	ItemsCount       *int   `query:"items_count,omitempty"`
}

// ListAddressLogsParams represents query parameters for listing address logs.
type ListAddressLogsParams struct {
	Topic      string `query:"topic,omitempty"`
	BlockNumber *int  `query:"block_number,omitempty"`
	Index       *int  `query:"index,omitempty"`
	ItemsCount  *int  `query:"items_count,omitempty"`
}

// ListAddressTokensParams represents query parameters for listing address tokens.
type ListAddressTokensParams struct {
	Type       string `query:"type,omitempty"` // comma-separated: ERC-20, ERC-721, ERC-1155, ERC-404
	FiatValue  string `query:"fiat_value,omitempty"`
	ID         *int   `query:"id,omitempty"`
	Value      string `query:"value,omitempty"`
	ItemsCount *int   `query:"items_count,omitempty"`
}

// ListAddressNFTsParams represents query parameters for listing address NFTs.
type ListAddressNFTsParams struct {
	Type                     string `query:"type,omitempty"` // ERC-721, ERC-1155, ERC-404
	TokenContractAddressHash string `query:"token_contract_address_hash,omitempty"`
	TokenID                  string `query:"token_id,omitempty"`
	TokenType                string `query:"token_type,omitempty"`
	ItemsCount               *int   `query:"items_count,omitempty"`
}

// ListAddressNFTCollectionsParams represents query parameters for listing address NFT collections.
type ListAddressNFTCollectionsParams struct {
	Type                     string `query:"type,omitempty"` // ERC-721, ERC-1155, ERC-404
	TokenContractAddressHash string `query:"token_contract_address_hash,omitempty"`
	TokenType                string `query:"token_type,omitempty"`
	ItemsCount               *int   `query:"items_count,omitempty"`
}

// ── RESPONSE TYPES ──

// AddressDetail represents detailed address information.
type AddressDetail struct {
	Address
	CoinBalance                string   `json:"coin_balance"`
	ExchangeRate               *string  `json:"exchange_rate"`
	BlockNumberBalanceUpdatedAt *int    `json:"block_number_balance_updated_at"`
	CreatorAddressHash         *string  `json:"creator_address_hash"`
	CreationTransactionHash    *string  `json:"creation_transaction_hash"`
	CreationStatus             *string  `json:"creation_status"`
	Token                      *Token   `json:"token"`
	HasValidatedBlocks         bool     `json:"has_validated_blocks"`
	HasLogs                    bool     `json:"has_logs"`
	HasTokens                  bool     `json:"has_tokens"`
	HasTokenTransfers          bool     `json:"has_token_transfers"`
	HasBeaconChainWithdrawals  bool     `json:"has_beacon_chain_withdrawals"`
	WatchlistAddressID         *int     `json:"watchlist_address_id"`
}

// AddressListItem represents an address in the addresses list.
type AddressListItem struct {
	Address
	CoinBalance       string  `json:"coin_balance"`
	ExchangeRate      *string `json:"exchange_rate"`
	TransactionsCount string  `json:"transactions_count"`
	TokenHoldingsCount *int   `json:"token_holdings_count"`
}

// AddressTokenBalance represents a token balance entry for an address.
type AddressTokenBalance struct {
	Token         Token          `json:"token"`
	Value         string         `json:"value"`
	TokenID       *string        `json:"token_id"`
	TokenInstance *TokenInstance  `json:"token_instance"`
}

// TokenBalance represents a token balance.
type TokenBalance struct {
	Value         string         `json:"value"`
	Token         *Token         `json:"token"`
	TokenID       *string        `json:"token_id"`
	TokenInstance *TokenInstance  `json:"token_instance"`
}

// TokenInstance represents an NFT token instance.
type TokenInstance struct {
	ID            string      `json:"id"`
	Metadata      interface{} `json:"metadata"`
	Owner         *Address    `json:"owner"`
	Token         *Token      `json:"token"`
	ExternalAppURL *string    `json:"external_app_url"`
	AnimationURL  *string     `json:"animation_url"`
	ImageURL      *string     `json:"image_url"`
	IsUnique      *bool       `json:"is_unique"`
	Thumbnails    interface{} `json:"thumbnails"`
	MediaType     *string     `json:"media_type"`
	MediaURL      *string     `json:"media_url"`
}

// NFTCollection represents a group of NFTs from the same collection.
type NFTCollection struct {
	Token          Token           `json:"token"`
	Amount         *string         `json:"amount"`
	TokenInstances []TokenInstance `json:"token_instances"`
}

// AddressCounters represents activity count stats for an address.
type AddressCounters struct {
	TransactionsCount  string `json:"transactions_count"`
	TokenTransfersCount string `json:"token_transfers_count"`
	GasUsageCount      string `json:"gas_usage_count"`
	ValidationsCount   string `json:"validations_count"`
}

// AddressTabsCounters represents tab counters for an address.
type AddressTabsCounters struct {
	TransactionsCount         int `json:"transactions_count"`
	InternalTransactionsCount int `json:"internal_transactions_count"`
	TokenTransfersCount       int `json:"token_transfers_count"`
	TokenBalancesCount        int `json:"token_balances_count"`
	LogsCount                 int `json:"logs_count"`
	ValidationsCount          int `json:"validations_count"`
	WithdrawalsCount          int `json:"withdrawals_count"`
}

// ── API METHODS ──

func (c *client) GetAddress(ctx context.Context, addressHash string) (resp *AddressDetail, err error) {
	path := buildPath("/api/v2/addresses/{address_hash}", map[string]string{
		"address_hash": addressHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) ListAddresses(ctx context.Context, params *ListAddressesParams) (resp *PaginatedResponse[AddressListItem], err error) {
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get("/api/v2/addresses")
	return
}

func (c *client) ListAddressTransactions(ctx context.Context, addressHash string, params *ListAddressTransactionsParams) (resp *PaginatedResponse[Transaction], err error) {
	path := buildPath("/api/v2/addresses/{address_hash}/transactions", map[string]string{
		"address_hash": addressHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) ListAddressTokenTransfers(ctx context.Context, addressHash string, params *ListAddressTokenTransfersParams) (resp *PaginatedResponse[TokenTransfer], err error) {
	path := buildPath("/api/v2/addresses/{address_hash}/token-transfers", map[string]string{
		"address_hash": addressHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) ListAddressInternalTransactions(ctx context.Context, addressHash string, params *ListAddressInternalTransactionsParams) (resp *PaginatedResponse[InternalTransaction], err error) {
	path := buildPath("/api/v2/addresses/{address_hash}/internal-transactions", map[string]string{
		"address_hash": addressHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) ListAddressLogs(ctx context.Context, addressHash string, params *ListAddressLogsParams) (resp *PaginatedResponse[Log], err error) {
	path := buildPath("/api/v2/addresses/{address_hash}/logs", map[string]string{
		"address_hash": addressHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) ListAddressTokens(ctx context.Context, addressHash string, params *ListAddressTokensParams) (resp *PaginatedResponse[AddressTokenBalance], err error) {
	path := buildPath("/api/v2/addresses/{address_hash}/tokens", map[string]string{
		"address_hash": addressHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) GetAddressTokenBalances(ctx context.Context, addressHash string) (resp []TokenBalance, err error) {
	path := buildPath("/api/v2/addresses/{address_hash}/token-balances", map[string]string{
		"address_hash": addressHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) ListAddressNFTs(ctx context.Context, addressHash string, params *ListAddressNFTsParams) (resp *PaginatedResponse[TokenInstance], err error) {
	path := buildPath("/api/v2/addresses/{address_hash}/nft", map[string]string{
		"address_hash": addressHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) ListAddressNFTCollections(ctx context.Context, addressHash string, params *ListAddressNFTCollectionsParams) (resp *PaginatedResponse[NFTCollection], err error) {
	path := buildPath("/api/v2/addresses/{address_hash}/nft/collections", map[string]string{
		"address_hash": addressHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) ListAddressBlocksValidated(ctx context.Context, addressHash string, params *PaginationParams) (resp *PaginatedResponse[Block], err error) {
	path := buildPath("/api/v2/addresses/{address_hash}/blocks-validated", map[string]string{
		"address_hash": addressHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) ListAddressCoinBalanceHistory(ctx context.Context, addressHash string, params *PaginationParams) (resp *PaginatedResponse[CoinBalance], err error) {
	path := buildPath("/api/v2/addresses/{address_hash}/coin-balance-history", map[string]string{
		"address_hash": addressHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) GetAddressCoinBalanceHistoryByDay(ctx context.Context, addressHash string) (resp []CoinBalanceByDay, err error) {
	path := buildPath("/api/v2/addresses/{address_hash}/coin-balance-history-by-day", map[string]string{
		"address_hash": addressHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) GetAddressCounters(ctx context.Context, addressHash string) (resp *AddressCounters, err error) {
	path := buildPath("/api/v2/addresses/{address_hash}/counters", map[string]string{
		"address_hash": addressHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) GetAddressTabsCounters(ctx context.Context, addressHash string) (resp *AddressTabsCounters, err error) {
	path := buildPath("/api/v2/addresses/{address_hash}/tabs-counters", map[string]string{
		"address_hash": addressHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) ListAddressWithdrawals(ctx context.Context, addressHash string, params *PaginationParams) (resp *PaginatedResponse[Withdrawal], err error) {
	path := buildPath("/api/v2/addresses/{address_hash}/withdrawals", map[string]string{
		"address_hash": addressHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}
