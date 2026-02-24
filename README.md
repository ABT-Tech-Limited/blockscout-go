# blockscout-go

Go client library for the [Blockscout](https://www.blockscout.com/) REST API v2.

## Install

```bash
go get github.com/ABT-Tech-Limited/blockscout-go
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	blockscout "github.com/ABT-Tech-Limited/blockscout-go"
)

func main() {
	// Create a client for any Blockscout instance
	client := blockscout.New("https://eth.blockscout.com", "")

	// Get address info
	addr, err := client.GetAddress(context.Background(), "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance: %s\n", addr.CoinBalance)
}
```

## Configuration

### With API Key

```go
client := blockscout.New("https://eth.blockscout.com", "your-api-key")
```

### With Options

```go
client := blockscout.New("https://eth.blockscout.com", "", blockscout.Options{
	Timeout: 15 * time.Second,
	Verbose: true, // enable debug logging
})
```

### With Rate Limiting

Blockscout free tier limits to ~5 requests per second. Use the built-in rate limiter:

```go
client := blockscout.New("https://eth.blockscout.com", "", blockscout.Options{
	BeforeRequest: []resty.RequestMiddleware{
		blockscout.FreeRateLimiter(),
	},
})
```

### With Custom HTTP Client

```go
restyCli := resty.New()
// configure restyCli as needed...
client := blockscout.NewWithClient("https://eth.blockscout.com", "", restyCli)
```

## API Coverage

### Addresses

```go
client.GetAddress(ctx, addressHash)
client.ListAddresses(ctx, params)
client.ListAddressTransactions(ctx, addressHash, params)
client.ListAddressTokenTransfers(ctx, addressHash, params)
client.ListAddressInternalTransactions(ctx, addressHash, params)
client.ListAddressLogs(ctx, addressHash, params)
client.ListAddressTokens(ctx, addressHash, params)
client.GetAddressTokenBalances(ctx, addressHash)
client.ListAddressNFTs(ctx, addressHash, params)
client.ListAddressNFTCollections(ctx, addressHash, params)
client.ListAddressBlocksValidated(ctx, addressHash, params)
client.ListAddressCoinBalanceHistory(ctx, addressHash, params)
client.GetAddressCoinBalanceHistoryByDay(ctx, addressHash)
client.GetAddressCounters(ctx, addressHash)
client.GetAddressTabsCounters(ctx, addressHash)
client.ListAddressWithdrawals(ctx, addressHash, params)
```

### Transactions

```go
client.ListTransactions(ctx, params)
client.GetTransaction(ctx, txHash)
client.GetTransactionStats(ctx)
client.ListTransactionTokenTransfers(ctx, txHash, params)
client.ListTransactionInternalTransactions(ctx, txHash, params)
client.ListTransactionLogs(ctx, txHash, params)
client.GetTransactionRawTrace(ctx, txHash)
client.ListTransactionStateChanges(ctx, txHash, params)
client.GetTransactionSummary(ctx, txHash)
```

### Blocks

```go
client.ListBlocks(ctx, params)
client.GetBlock(ctx, blockHashOrNumber)
client.ListBlockTransactions(ctx, blockHashOrNumber, params)
client.ListBlockWithdrawals(ctx, blockHashOrNumber, params)
```

### Tokens

```go
client.ListTokens(ctx, params)
client.GetToken(ctx, addressHash)
client.ListTokenTransfers(ctx, addressHash, params)
client.ListTokenHolders(ctx, addressHash, params)
client.GetTokenCounters(ctx, addressHash)
client.ListTokenInstances(ctx, addressHash, params)
client.GetTokenInstance(ctx, addressHash, tokenID)
client.ListTokenInstanceTransfers(ctx, addressHash, tokenID, params)
client.ListTokenInstanceHolders(ctx, addressHash, tokenID, params)
```

### Smart Contracts

```go
client.ListSmartContracts(ctx, params)
client.GetSmartContract(ctx, addressHash)
client.GetSmartContractCounters(ctx)
client.GetSmartContractReadMethods(ctx, addressHash)
client.GetSmartContractWriteMethods(ctx, addressHash)
client.GetSmartContractReadMethodsProxy(ctx, addressHash)
client.GetSmartContractWriteMethodsProxy(ctx, addressHash)
client.QuerySmartContractReadMethod(ctx, addressHash, req)
```

### Stats

```go
client.GetStats(ctx)
client.GetTransactionsChart(ctx)
client.GetMarketChart(ctx)
```

### Search

```go
client.Search(ctx, params)
client.CheckRedirect(ctx, query)
```

## Pagination

List endpoints return `*PaginatedResponse[T]` with cursor-based pagination:

```go
// First page
resp, err := client.ListTransactions(ctx, &blockscout.ListTransactionsParams{
	Sort:  "block_number",
	Order: "desc",
})

fmt.Println(resp.Items)           // []Transaction
fmt.Println(resp.HasNextPage())   // true if more pages exist
fmt.Println(resp.NextPageParams)  // pass to next request for cursor pagination
```

## Debug Mode

Enable verbose HTTP logging for troubleshooting:

```go
debugClient := client.Debug()
debugClient.GetAddress(ctx, "0x...")
```

## License

MIT
