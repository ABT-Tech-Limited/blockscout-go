package blockscout

import (
	"os"
	"time"

	"resty.dev/v3"
)

var (
	api     Client
	baseURL = "https://teleskop.aeris-dev.codedsolution-web3.com"
)

func init() {
	if envURL := os.Getenv("BLOCKSCOUT_BASE_URL"); envURL != "" {
		baseURL = envURL
	}

	apiKey := os.Getenv("BLOCKSCOUT_API_KEY")

	api = New(baseURL, apiKey, Options{
		Timeout: time.Second * 15,
		Verbose: false,
		BeforeRequest: []resty.RequestMiddleware{
			FreeRateLimiter(),
		},
	})
}
