package blockscout

import "context"

// ── RESPONSE TYPES ──

// NetworkStats represents overall network statistics.
type NetworkStats struct {
	AverageBlockTime             float64   `json:"average_block_time"`
	CoinImage                    *string   `json:"coin_image"`
	CoinPrice                    *string   `json:"coin_price"`
	CoinPriceChangePercentage    *float64  `json:"coin_price_change_percentage"`
	GasPriceUpdatedAt            string    `json:"gas_price_updated_at"`
	GasPrices                    GasPrices `json:"gas_prices"`
	GasPricesUpdateIn            int       `json:"gas_prices_update_in"`
	GasUsedToday                 string    `json:"gas_used_today"`
	MarketCap                    string    `json:"market_cap"`
	NetworkUtilizationPercentage float64   `json:"network_utilization_percentage"`
	SecondaryCoinImage           *string   `json:"secondary_coin_image"`
	SecondaryCoinPrice           *string   `json:"secondary_coin_price"`
	StaticGasPrice               *string   `json:"static_gas_price"`
	TotalAddresses               string    `json:"total_addresses"`
	TotalBlocks                  string    `json:"total_blocks"`
	TotalGasUsed                 string    `json:"total_gas_used"`
	TotalTransactions            string    `json:"total_transactions"`
	TransactionsToday            string    `json:"transactions_today"`
	TVL                          *string   `json:"tvl"`
}

// GasPrices represents gas price recommendations.
type GasPrices struct {
	Slow    *float64 `json:"slow"`
	Average *float64 `json:"average"`
	Fast    *float64 `json:"fast"`
}

// TransactionsChart represents daily transaction count chart data.
type TransactionsChart struct {
	ChartData []TransactionsChartItem `json:"chart_data"`
}

// TransactionsChartItem represents a single data point in the transactions chart.
type TransactionsChartItem struct {
	Date              string `json:"date"`
	TransactionsCount int    `json:"transactions_count"`
}

// MarketChart represents market chart data.
type MarketChart struct {
	ChartData []MarketChartItem `json:"chart_data"`
}

// MarketChartItem represents a single data point in the market chart.
type MarketChartItem struct {
	Date                    string  `json:"date"`
	ClosingPrice            *string `json:"closing_price"`
	MarketCap               *string `json:"market_cap"`
	TVL                     *string `json:"tvl"`
	SecondaryCoinClosingPrice *string `json:"secondary_coin_closing_price"`
	SecondaryCoinMarketCap    *string `json:"secondary_coin_market_cap"`
}

// ── API METHODS ──

func (c *client) GetStats(ctx context.Context) (resp *NetworkStats, err error) {
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get("/api/v2/stats")
	return
}

func (c *client) GetTransactionsChart(ctx context.Context) (resp *TransactionsChart, err error) {
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get("/api/v2/stats/charts/transactions")
	return
}

func (c *client) GetMarketChart(ctx context.Context) (resp *MarketChart, err error) {
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get("/api/v2/stats/charts/market")
	return
}
