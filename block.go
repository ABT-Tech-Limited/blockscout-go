package blockscout

import "context"

// ── REQUEST TYPES ──

// ListBlocksParams represents query parameters for listing blocks.
type ListBlocksParams struct {
	Type        string `query:"type,omitempty"` // "block" | "uncle" | "reorg"
	BlockNumber *int   `query:"block_number,omitempty"`
	ItemsCount  *int   `query:"items_count,omitempty"`
}

// ── RESPONSE TYPES ──

// Block represents a block object.
type Block struct {
	Height                    int       `json:"height"`
	Timestamp                 string    `json:"timestamp"`
	Hash                      string    `json:"hash"`
	ParentHash                string    `json:"parent_hash"`
	TransactionsCount         int       `json:"transactions_count"`
	InternalTransactionsCount *int      `json:"internal_transactions_count"`
	Miner                     Address   `json:"miner"`
	Size                      int       `json:"size"`
	Difficulty                *string   `json:"difficulty"`
	TotalDifficulty           *string   `json:"total_difficulty"`
	GasUsed                   string    `json:"gas_used"`
	GasLimit                  string    `json:"gas_limit"`
	GasUsedPercentage         float64   `json:"gas_used_percentage"`
	GasTargetPercentage       float64   `json:"gas_target_percentage"`
	Nonce                     *string   `json:"nonce"`
	BaseFeePerGas             *string   `json:"base_fee_per_gas"`
	BurntFees                 *string   `json:"burnt_fees"`
	BurntFeesPercentage       *float64  `json:"burnt_fees_percentage"`
	PriorityFee               *string   `json:"priority_fee"`
	TransactionFees           string    `json:"transaction_fees"`
	Type                      string    `json:"type"`
	Rewards                   []BlockReward  `json:"rewards"`
	UnclesHashes              []BlockUncle   `json:"uncles_hashes"`
	WithdrawalsCount          *int      `json:"withdrawals_count"`
	BlobTransactionsCount     *int      `json:"blob_transactions_count"`
	BlobGasUsed               *string   `json:"blob_gas_used"`
	ExcessBlobGas             *string   `json:"excess_blob_gas"`
}

// BlockReward represents a block reward entry.
type BlockReward struct {
	Type   string `json:"type"`
	Reward string `json:"reward"`
}

// BlockUncle represents an uncle block hash.
type BlockUncle struct {
	Hash string `json:"hash"`
}

// ── API METHODS ──

func (c *client) ListBlocks(ctx context.Context, params *ListBlocksParams) (resp *PaginatedResponse[Block], err error) {
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get("/api/v2/blocks")
	return
}

func (c *client) GetBlock(ctx context.Context, blockHashOrNumber string) (resp *Block, err error) {
	path := buildPath("/api/v2/blocks/{id}", map[string]string{
		"id": blockHashOrNumber,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) ListBlockTransactions(ctx context.Context, blockHashOrNumber string, params *PaginationParams) (resp *PaginatedResponse[Transaction], err error) {
	path := buildPath("/api/v2/blocks/{id}/transactions", map[string]string{
		"id": blockHashOrNumber,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) ListBlockWithdrawals(ctx context.Context, blockHashOrNumber string, params *PaginationParams) (resp *PaginatedResponse[Withdrawal], err error) {
	path := buildPath("/api/v2/blocks/{id}/withdrawals", map[string]string{
		"id": blockHashOrNumber,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}
