package blockscout

import (
	"context"
	"encoding/json"
)

// ── REQUEST TYPES ──

// ListTransactionsParams represents query parameters for listing transactions.
type ListTransactionsParams struct {
	Filter      string `query:"filter,omitempty"`
	Type        string `query:"type,omitempty"`
	Method      string `query:"method,omitempty"`
	BlockNumber *int   `query:"block_number,omitempty"`
	Index       *int   `query:"index,omitempty"`
	ItemsCount  *int   `query:"items_count,omitempty"`
}

// ── RESPONSE TYPES ──

// Transaction represents a transaction object.
type Transaction struct {
	Hash                            string           `json:"hash"`
	Result                          string           `json:"result"`
	Status                          *string          `json:"status"`
	BlockNumber                     *int             `json:"block_number"`
	Timestamp                       *string          `json:"timestamp"`
	From                            Address          `json:"from"`
	To                              *Address         `json:"to"`
	CreatedContract                 *Address         `json:"created_contract"`
	Confirmations                   int              `json:"confirmations"`
	ConfirmationDuration            []float64        `json:"confirmation_duration"`
	Value                           string           `json:"value"`
	Fee                             Fee              `json:"fee"`
	GasPrice                        *string          `json:"gas_price"`
	Type                            *int             `json:"type"`
	GasUsed                         *string          `json:"gas_used"`
	GasLimit                        string           `json:"gas_limit"`
	MaxFeePerGas                    *string          `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas            *string          `json:"max_priority_fee_per_gas"`
	BaseFeePerGas                   *string          `json:"base_fee_per_gas"`
	PriorityFee                     *string          `json:"priority_fee"`
	TransactionBurntFee             *string          `json:"transaction_burnt_fee"`
	Nonce                           int              `json:"nonce"`
	Position                        *int             `json:"position"`
	RevertReason                    json.RawMessage  `json:"revert_reason"`
	RawInput                        string           `json:"raw_input"`
	DecodedInput                    *DecodedInput    `json:"decoded_input"`
	TokenTransfers                  []TokenTransfer  `json:"token_transfers"`
	TokenTransfersOverflow          *bool            `json:"token_transfers_overflow"`
	Actions                         []TransactionAction `json:"actions"`
	ExchangeRate                    *string          `json:"exchange_rate"`
	HistoricExchangeRate            *string          `json:"historic_exchange_rate"`
	Method                          *string          `json:"method"`
	TransactionTypes                []string         `json:"transaction_types"`
	TransactionTag                  *string          `json:"transaction_tag"`
	HasErrorInInternalTransactions  *bool            `json:"has_error_in_internal_transactions"`
	AuthorizationList               []SignedAuthorization `json:"authorization_list"`
}

// TransactionAction represents an action performed during a transaction.
type TransactionAction struct {
	Protocol string      `json:"protocol"`
	Type     string      `json:"type"`
	Data     interface{} `json:"data"`
}

// SignedAuthorization represents a signed authorization (EIP-7702).
type SignedAuthorization struct {
	AddressHash string `json:"address_hash"`
	Authority   string `json:"authority"`
	ChainID     int    `json:"chain_id"`
	Nonce       string `json:"nonce"`
	R           string `json:"r"`
	S           string `json:"s"`
	V           int    `json:"v"`
}

// TransactionStats represents transaction statistics.
type TransactionStats struct {
	TransactionsCount24h    string `json:"transactions_count_24h"`
	PendingTransactionsCount string `json:"pending_transactions_count"`
	TransactionFeesSum24h   string `json:"transaction_fees_sum_24h"`
	TransactionFeesAvg24h   string `json:"transaction_fees_avg_24h"`
}

// StateChange represents a state change caused by a transaction.
type StateChange struct {
	Address       Address          `json:"address"`
	IsMinor       bool             `json:"is_miner"`
	BalanceBefore *string          `json:"balance_before"`
	BalanceAfter  *string          `json:"balance_after"`
	TokenBalancesBefore []interface{} `json:"token_balances_before"`
	TokenBalancesAfter  []interface{} `json:"token_balances_after"`
}

// TransactionSummary represents a human-readable transaction summary.
type TransactionSummary struct {
	Result string `json:"result"`
}

// ── API METHODS ──

func (c *client) ListTransactions(ctx context.Context, params *ListTransactionsParams) (resp *PaginatedResponse[Transaction], err error) {
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get("/api/v2/transactions")
	return
}

func (c *client) GetTransaction(ctx context.Context, txHash string) (resp *Transaction, err error) {
	path := buildPath("/api/v2/transactions/{tx_hash}", map[string]string{
		"tx_hash": txHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) GetTransactionStats(ctx context.Context) (resp *TransactionStats, err error) {
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get("/api/v2/transactions/stats")
	return
}

func (c *client) ListTransactionTokenTransfers(ctx context.Context, txHash string, params *PaginationParams) (resp *PaginatedResponse[TokenTransfer], err error) {
	path := buildPath("/api/v2/transactions/{tx_hash}/token-transfers", map[string]string{
		"tx_hash": txHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) ListTransactionInternalTransactions(ctx context.Context, txHash string, params *PaginationParams) (resp *PaginatedResponse[InternalTransaction], err error) {
	path := buildPath("/api/v2/transactions/{tx_hash}/internal-transactions", map[string]string{
		"tx_hash": txHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) ListTransactionLogs(ctx context.Context, txHash string, params *PaginationParams) (resp *PaginatedResponse[Log], err error) {
	path := buildPath("/api/v2/transactions/{tx_hash}/logs", map[string]string{
		"tx_hash": txHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) GetTransactionRawTrace(ctx context.Context, txHash string) (resp []RawTrace, err error) {
	path := buildPath("/api/v2/transactions/{tx_hash}/raw-trace", map[string]string{
		"tx_hash": txHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) ListTransactionStateChanges(ctx context.Context, txHash string, params *PaginationParams) (resp *PaginatedResponse[StateChange], err error) {
	path := buildPath("/api/v2/transactions/{tx_hash}/state-changes", map[string]string{
		"tx_hash": txHash,
	})
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get(path)
	return
}

func (c *client) GetTransactionSummary(ctx context.Context, txHash string) (resp *TransactionSummary, err error) {
	path := buildPath("/api/v2/transactions/{tx_hash}/summary", map[string]string{
		"tx_hash": txHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}
