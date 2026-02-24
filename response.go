package blockscout

import "encoding/json"

// PaginatedResponse represents a paginated API response.
type PaginatedResponse[T any] struct {
	Items          []T                    `json:"items"`
	NextPageParams map[string]interface{} `json:"next_page_params"`
}

// HasNextPage returns true if there are more pages available.
func (r *PaginatedResponse[T]) HasNextPage() bool {
	return r != nil && r.NextPageParams != nil && len(r.NextPageParams) > 0
}

// PaginationParams provides common pagination query parameters.
type PaginationParams struct {
	BlockNumber *int `query:"block_number,omitempty"`
	Index       *int `query:"index,omitempty"`
	ItemsCount  *int `query:"items_count,omitempty"`
}

// Address represents a common address object shared across all endpoints.
type Address struct {
	Hash           string           `json:"hash"`
	IsContract     *bool            `json:"is_contract"`
	IsScam         bool             `json:"is_scam"`
	IsVerified     bool             `json:"is_verified"`
	Name           *string          `json:"name"`
	ENSDomainName  *string          `json:"ens_domain_name"`
	ProxyType      *string          `json:"proxy_type"`
	Implementations []Implementation `json:"implementations"`
	Metadata       *Metadata        `json:"metadata"`
	PrivateTags    []Tag            `json:"private_tags"`
	PublicTags     []Tag            `json:"public_tags"`
	WatchlistNames []WatchlistName  `json:"watchlist_names"`
}

// Implementation represents a proxy implementation.
type Implementation struct {
	AddressHash string  `json:"address_hash"`
	Name        *string `json:"name"`
}

// Metadata contains address metadata tags.
type Metadata struct {
	Tags []MetadataTag `json:"tags"`
}

// MetadataTag represents a metadata tag.
type MetadataTag struct {
	Slug    string      `json:"slug"`
	Name    string      `json:"name"`
	TagType string      `json:"tagType"`
	Ordinal int         `json:"ordinal"`
	Meta    interface{} `json:"meta"`
}

// Tag represents an address tag (private or public).
type Tag struct {
	AddressHash string `json:"address_hash"`
	DisplayName string `json:"display_name"`
	Label       string `json:"label"`
}

// WatchlistName represents a watchlist entry.
type WatchlistName struct {
	DisplayName string `json:"display_name"`
	Label       string `json:"label"`
}

// Token represents a token object.
type Token struct {
	AddressHash         string  `json:"address_hash"`
	Name                *string `json:"name"`
	Symbol              *string `json:"symbol"`
	Decimals            *string `json:"decimals"`
	Type                *string `json:"type"`
	TotalSupply         *string `json:"total_supply"`
	HoldersCount        *string `json:"holders_count"`
	ExchangeRate        *string `json:"exchange_rate"`
	CirculatingMarketCap *string `json:"circulating_market_cap"`
	Volume24h           *string `json:"volume_24h"`
	IconURL             *string `json:"icon_url"`
}

// Fee represents a transaction fee.
type Fee struct {
	Type  string  `json:"type"`
	Value *string `json:"value"`
}

// DecodedInput represents decoded transaction input data.
type DecodedInput struct {
	MethodID   *string              `json:"method_id"`
	MethodCall *string              `json:"method_call"`
	Parameters []DecodedInputParam  `json:"parameters"`
}

// DecodedInputParam represents a parameter in decoded input.
type DecodedInputParam struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

// TokenTransfer represents a token transfer event.
type TokenTransfer struct {
	TransactionHash string       `json:"transaction_hash"`
	BlockHash       string       `json:"block_hash"`
	BlockNumber     int          `json:"block_number"`
	Timestamp       *string      `json:"timestamp"`
	From            Address      `json:"from"`
	To              Address      `json:"to"`
	Token           Token        `json:"token"`
	Total           *TokenTotal  `json:"total"`
	Type            string       `json:"type"`
	Method          *string      `json:"method"`
	LogIndex        int          `json:"log_index"`
}

// TokenTotal represents the value of a token transfer.
type TokenTotal struct {
	Value         *string        `json:"value"`
	Decimals      *string        `json:"decimals"`
	TokenID       *string        `json:"token_id"`
	TokenInstance *TokenInstance  `json:"token_instance"`
}

// InternalTransaction represents an internal transaction.
type InternalTransaction struct {
	TransactionHash  string  `json:"transaction_hash"`
	BlockNumber      int     `json:"block_number"`
	Timestamp        string  `json:"timestamp"`
	From             Address `json:"from"`
	To               Address `json:"to"`
	CreatedContract  *Address `json:"created_contract"`
	Value            string  `json:"value"`
	Type             string  `json:"type"`
	Success          bool    `json:"success"`
	Error            *string `json:"error"`
	GasLimit         *string `json:"gas_limit"`
	Index            int     `json:"index"`
	BlockIndex       int     `json:"block_index"`
	TransactionIndex int     `json:"transaction_index"`
}

// Log represents an event log.
type Log struct {
	TransactionHash string       `json:"transaction_hash"`
	Address         Address      `json:"address"`
	Topics          []string     `json:"topics"`
	Data            string       `json:"data"`
	Index           int          `json:"index"`
	Decoded         *DecodedInput `json:"decoded"`
	SmartContract   *Address     `json:"smart_contract"`
	BlockHash       string       `json:"block_hash"`
	BlockNumber     int          `json:"block_number"`
}

// Withdrawal represents a validator withdrawal.
type Withdrawal struct {
	Index          int      `json:"index"`
	ValidatorIndex int      `json:"validator_index"`
	Amount         string   `json:"amount"`
	BlockNumber    *int     `json:"block_number"`
	Receiver       *Address `json:"receiver"`
	Timestamp      *string  `json:"timestamp"`
}

// CoinBalance represents a native coin balance entry.
type CoinBalance struct {
	TransactionHash *string `json:"transaction_hash"`
	BlockNumber     int     `json:"block_number"`
	BlockTimestamp  string  `json:"block_timestamp"`
	Delta           string  `json:"delta"`
	Value           string  `json:"value"`
}

// CoinBalanceByDay represents daily coin balance snapshot.
type CoinBalanceByDay struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}

// RawTrace represents a raw execution trace entry.
type RawTrace = json.RawMessage
