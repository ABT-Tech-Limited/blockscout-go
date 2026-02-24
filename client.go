package blockscout

import (
	"context"
	"net/http"
	"strings"
	"time"

	"golang.org/x/time/rate"
	"resty.dev/v3"
)

// Client defines the interface for interacting with the Blockscout REST API v2.
type Client interface {
	// Address endpoints
	GetAddress(ctx context.Context, addressHash string) (*AddressDetail, error)
	ListAddresses(ctx context.Context, params *ListAddressesParams) (*PaginatedResponse[AddressListItem], error)
	ListAddressTransactions(ctx context.Context, addressHash string, params *ListAddressTransactionsParams) (*PaginatedResponse[Transaction], error)
	ListAddressTokenTransfers(ctx context.Context, addressHash string, params *ListAddressTokenTransfersParams) (*PaginatedResponse[TokenTransfer], error)
	ListAddressInternalTransactions(ctx context.Context, addressHash string, params *ListAddressInternalTransactionsParams) (*PaginatedResponse[InternalTransaction], error)
	ListAddressLogs(ctx context.Context, addressHash string, params *ListAddressLogsParams) (*PaginatedResponse[Log], error)
	ListAddressTokens(ctx context.Context, addressHash string, params *ListAddressTokensParams) (*PaginatedResponse[AddressTokenBalance], error)
	GetAddressTokenBalances(ctx context.Context, addressHash string) ([]TokenBalance, error)
	ListAddressNFTs(ctx context.Context, addressHash string, params *ListAddressNFTsParams) (*PaginatedResponse[TokenInstance], error)
	ListAddressNFTCollections(ctx context.Context, addressHash string, params *ListAddressNFTCollectionsParams) (*PaginatedResponse[NFTCollection], error)
	ListAddressBlocksValidated(ctx context.Context, addressHash string, params *PaginationParams) (*PaginatedResponse[Block], error)
	ListAddressCoinBalanceHistory(ctx context.Context, addressHash string, params *PaginationParams) (*PaginatedResponse[CoinBalance], error)
	GetAddressCoinBalanceHistoryByDay(ctx context.Context, addressHash string) ([]CoinBalanceByDay, error)
	GetAddressCounters(ctx context.Context, addressHash string) (*AddressCounters, error)
	GetAddressTabsCounters(ctx context.Context, addressHash string) (*AddressTabsCounters, error)
	ListAddressWithdrawals(ctx context.Context, addressHash string, params *PaginationParams) (*PaginatedResponse[Withdrawal], error)

	// Transaction endpoints
	ListTransactions(ctx context.Context, params *ListTransactionsParams) (*PaginatedResponse[Transaction], error)
	GetTransaction(ctx context.Context, txHash string) (*Transaction, error)
	GetTransactionStats(ctx context.Context) (*TransactionStats, error)
	ListTransactionTokenTransfers(ctx context.Context, txHash string, params *PaginationParams) (*PaginatedResponse[TokenTransfer], error)
	ListTransactionInternalTransactions(ctx context.Context, txHash string, params *PaginationParams) (*PaginatedResponse[InternalTransaction], error)
	ListTransactionLogs(ctx context.Context, txHash string, params *PaginationParams) (*PaginatedResponse[Log], error)
	GetTransactionRawTrace(ctx context.Context, txHash string) ([]RawTrace, error)
	ListTransactionStateChanges(ctx context.Context, txHash string, params *PaginationParams) (*PaginatedResponse[StateChange], error)
	GetTransactionSummary(ctx context.Context, txHash string) (*TransactionSummary, error)

	// Block endpoints
	ListBlocks(ctx context.Context, params *ListBlocksParams) (*PaginatedResponse[Block], error)
	GetBlock(ctx context.Context, blockHashOrNumber string) (*Block, error)
	ListBlockTransactions(ctx context.Context, blockHashOrNumber string, params *PaginationParams) (*PaginatedResponse[Transaction], error)
	ListBlockWithdrawals(ctx context.Context, blockHashOrNumber string, params *PaginationParams) (*PaginatedResponse[Withdrawal], error)

	// Token endpoints
	ListTokens(ctx context.Context, params *ListTokensParams) (*PaginatedResponse[Token], error)
	GetToken(ctx context.Context, addressHash string) (*Token, error)
	ListTokenTransfers(ctx context.Context, addressHash string, params *PaginationParams) (*PaginatedResponse[TokenTransfer], error)
	ListTokenHolders(ctx context.Context, addressHash string, params *PaginationParams) (*PaginatedResponse[TokenHolder], error)
	GetTokenCounters(ctx context.Context, addressHash string) (*TokenCounters, error)
	ListTokenInstances(ctx context.Context, addressHash string, params *PaginationParams) (*PaginatedResponse[TokenInstance], error)
	GetTokenInstance(ctx context.Context, addressHash string, tokenID string) (*TokenInstance, error)
	ListTokenInstanceTransfers(ctx context.Context, addressHash string, tokenID string, params *PaginationParams) (*PaginatedResponse[TokenTransfer], error)
	ListTokenInstanceHolders(ctx context.Context, addressHash string, tokenID string, params *PaginationParams) (*PaginatedResponse[TokenHolder], error)

	// Smart Contract endpoints
	ListSmartContracts(ctx context.Context, params *ListSmartContractsParams) (*PaginatedResponse[SmartContractListItem], error)
	GetSmartContract(ctx context.Context, addressHash string) (*SmartContract, error)
	GetSmartContractCounters(ctx context.Context) (*SmartContractCounters, error)
	GetSmartContractReadMethods(ctx context.Context, addressHash string) ([]SmartContractMethod, error)
	GetSmartContractWriteMethods(ctx context.Context, addressHash string) ([]SmartContractMethod, error)
	GetSmartContractReadMethodsProxy(ctx context.Context, addressHash string) ([]SmartContractMethod, error)
	GetSmartContractWriteMethodsProxy(ctx context.Context, addressHash string) ([]SmartContractMethod, error)
	QuerySmartContractReadMethod(ctx context.Context, addressHash string, req QueryReadMethodReq) ([]SmartContractReadResult, error)

	// Stats endpoints
	GetStats(ctx context.Context) (*NetworkStats, error)
	GetTransactionsChart(ctx context.Context) (*TransactionsChart, error)
	GetMarketChart(ctx context.Context) (*MarketChart, error)

	// Search endpoints
	Search(ctx context.Context, params *SearchParams) (*PaginatedResponse[SearchResult], error)
	CheckRedirect(ctx context.Context, query string) (*SearchRedirect, error)

	// Debug returns a client with verbose logging enabled.
	Debug() Client
}

var _ Client = (*client)(nil)

type client struct {
	resty  *resty.Client
	apiKey string
}

// Options configures the Blockscout client behavior.
type Options struct {
	Timeout       time.Duration
	Verbose       bool
	Transport     *http.Transport
	BeforeRequest []resty.RequestMiddleware
}

// New creates a new Blockscout API client.
// baseURL is the Blockscout instance URL (e.g., "https://eth.blockscout.com").
// apiKey is optional; pass an empty string if not needed.
func New(baseURL string, apiKey string, opts ...Options) Client {
	baseURL = strings.TrimRight(baseURL, "/")

	restyCli := resty.New().
		SetBaseURL(baseURL).
		SetHeader("User-Agent", "blockscout-go-client").
		SetHeader("Accept", "application/json")

	if len(opts) > 0 {
		opt := opts[0]
		if opt.Timeout > 0 {
			restyCli.SetTimeout(opt.Timeout)
		} else {
			restyCli.SetTimeout(10 * time.Second)
		}
		if opt.Verbose {
			restyCli.SetDebug(true)
		}
		if opt.Transport != nil {
			restyCli.SetTransport(opt.Transport)
		}
		for _, mw := range opt.BeforeRequest {
			restyCli.AddRequestMiddleware(mw)
		}
	} else {
		restyCli.SetTimeout(10 * time.Second)
	}

	if apiKey != "" {
		restyCli.AddRequestMiddleware(func(c *resty.Client, req *resty.Request) error {
			req.SetQueryParam("apikey", apiKey)
			return nil
		})
	}

	return &client{
		resty:  restyCli,
		apiKey: apiKey,
	}
}

// NewWithClient creates a client with a pre-configured resty client.
func NewWithClient(baseURL string, apiKey string, restyCli *resty.Client) Client {
	baseURL = strings.TrimRight(baseURL, "/")
	restyCli.SetBaseURL(baseURL)

	if apiKey != "" {
		restyCli.AddRequestMiddleware(func(c *resty.Client, req *resty.Request) error {
			req.SetQueryParam("apikey", apiKey)
			return nil
		})
	}

	return &client{
		resty:  restyCli,
		apiKey: apiKey,
	}
}

// FreeRateLimiter returns a resty.RequestMiddleware that limits requests to 5 per second.
func FreeRateLimiter() resty.RequestMiddleware {
	limiter := rate.NewLimiter(rate.Every(time.Second/5), 1)
	return func(c *resty.Client, req *resty.Request) error {
		if err := limiter.Wait(req.Context()); err != nil {
			return err
		}
		return nil
	}
}

func (c *client) Debug() Client {
	return &client{
		resty:  c.resty.EnableDebug(),
		apiKey: c.apiKey,
	}
}
