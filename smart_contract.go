package blockscout

import "context"

// ── REQUEST TYPES ──

// ListSmartContractsParams represents query parameters for listing smart contracts.
type ListSmartContractsParams struct {
	Q               string `query:"q,omitempty"`                // search query
	Filter          string `query:"filter,omitempty"`           // "vyper" | "solidity" | "yul"
	ItemsCount      *int   `query:"items_count,omitempty"`
	SmartContractID *int   `query:"smart_contract_id,omitempty"`
}

// QueryReadMethodReq represents a request to query a smart contract read method.
type QueryReadMethodReq struct {
	Args         []interface{} `json:"args"`
	MethodID     string        `json:"method_id"`
	ContractType string        `json:"contract_type"`
	From         string        `json:"from,omitempty"`
}

// ── RESPONSE TYPES ──

// SmartContractListItem represents a smart contract in the list response.
type SmartContractListItem struct {
	Address             Address `json:"address"`
	Certified           bool    `json:"certified"`
	CoinBalance         string  `json:"coin_balance"`
	CompilerVersion     string  `json:"compiler_version"`
	HasConstructorArgs  bool    `json:"has_constructor_args"`
	Language            string  `json:"language"` // "solidity" | "vyper" | "yul" | "geas"
	LicenseType         string  `json:"license_type"`
	MarketCap           *string `json:"market_cap"`
	OptimizationEnabled bool    `json:"optimization_enabled"`
	TransactionsCount   int     `json:"transactions_count"`
	VerifiedAt          string  `json:"verified_at"`
}

// SmartContract represents detailed smart contract information.
type SmartContract struct {
	SmartContractListItem
	Name                            *string                `json:"name"`
	SourceCode                      *string                `json:"source_code"`
	DeployedBytecode                *string                `json:"deployed_bytecode"`
	CreationBytecode                *string                `json:"creation_bytecode"`
	ABI                             interface{}            `json:"abi"`
	ConstructorArgs                 *string                `json:"constructor_args"`
	DecodedConstructorArgs          interface{}            `json:"decoded_constructor_args"`
	OptimizationRuns                *int                   `json:"optimization_runs"`
	EVMVersion                      *string                `json:"evm_version"`
	CompilerSettings                interface{}            `json:"compiler_settings"`
	AdditionalSources               []AdditionalSource     `json:"additional_sources"`
	ExternalLibraries               []ExternalLibrary      `json:"external_libraries"`
	Implementations                 []Implementation       `json:"implementations"`
	ProxyType                       *string                `json:"proxy_type"`
	IsVerified                      *bool                  `json:"is_verified"`
	IsFullyVerified                 *bool                  `json:"is_fully_verified"`
	IsPartiallyVerified             *bool                  `json:"is_partially_verified"`
	IsVerifiedViaSourcify           *bool                  `json:"is_verified_via_sourcify"`
	IsVerifiedViaEthBytecodeDB      *bool                  `json:"is_verified_via_eth_bytecode_db"`
	IsVerifiedViaVerifierAlliance   *bool                  `json:"is_verified_via_verifier_alliance"`
	IsBlueprint                     *bool                  `json:"is_blueprint"`
	IsChangedBytecode               *bool                  `json:"is_changed_bytecode"`
	CanBeVisualizedViaSol2uml       *bool                  `json:"can_be_visualized_via_sol2uml"`
	FilePath                        *string                `json:"file_path"`
	SourcifyRepoURL                 *string                `json:"sourcify_repo_url"`
	VerifiedTwinAddressHash         *string                `json:"verified_twin_address_hash"`
}

// AdditionalSource represents an additional source file of a smart contract.
type AdditionalSource struct {
	FilePath   string `json:"file_path"`
	SourceCode string `json:"source_code"`
}

// ExternalLibrary represents an external library used by a smart contract.
type ExternalLibrary struct {
	AddressHash string `json:"address_hash"`
	Name        string `json:"name"`
}

// SmartContractCounters represents smart contract statistics.
type SmartContractCounters struct {
	SmartContracts                string `json:"smart_contracts"`
	VerifiedSmartContracts        string `json:"verified_smart_contracts"`
	NewSmartContracts24h          string `json:"new_smart_contracts_24h"`
	NewVerifiedSmartContracts24h  string `json:"new_verified_smart_contracts_24h"`
}

// SmartContractMethod represents a smart contract method (read or write).
type SmartContractMethod struct {
	Type            string                    `json:"type"`
	StateMutability string                    `json:"stateMutability"`
	Name            string                    `json:"name"`
	Inputs          []SmartContractMethodParam `json:"inputs"`
	Outputs         []SmartContractMethodParam `json:"outputs"`
	MethodID        string                    `json:"method_id"`
}

// SmartContractMethodParam represents a method input/output parameter.
type SmartContractMethodParam struct {
	InternalType string                    `json:"internalType"`
	Name         string                    `json:"name"`
	Type         string                    `json:"type"`
	Components   []SmartContractMethodParam `json:"components,omitempty"`
}

// SmartContractReadResult represents the result of querying a read method.
type SmartContractReadResult struct {
	IsError bool        `json:"is_error"`
	Result  interface{} `json:"result"`
}

// ── API METHODS ──

func (c *client) ListSmartContracts(ctx context.Context, params *ListSmartContractsParams) (resp *PaginatedResponse[SmartContractListItem], err error) {
	req := c.resty.R().SetContext(ctx).SetResult(&resp)
	if params != nil {
		req.SetQueryParams(structToQueryParams(params))
	}
	_, err = req.Get("/api/v2/smart-contracts")
	return
}

func (c *client) GetSmartContract(ctx context.Context, addressHash string) (resp *SmartContract, err error) {
	path := buildPath("/api/v2/smart-contracts/{address_hash}", map[string]string{
		"address_hash": addressHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) GetSmartContractCounters(ctx context.Context) (resp *SmartContractCounters, err error) {
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get("/api/v2/smart-contracts/counters")
	return
}

func (c *client) GetSmartContractReadMethods(ctx context.Context, addressHash string) (resp []SmartContractMethod, err error) {
	path := buildPath("/api/v2/smart-contracts/{address_hash}/methods-read", map[string]string{
		"address_hash": addressHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) GetSmartContractWriteMethods(ctx context.Context, addressHash string) (resp []SmartContractMethod, err error) {
	path := buildPath("/api/v2/smart-contracts/{address_hash}/methods-write", map[string]string{
		"address_hash": addressHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) GetSmartContractReadMethodsProxy(ctx context.Context, addressHash string) (resp []SmartContractMethod, err error) {
	path := buildPath("/api/v2/smart-contracts/{address_hash}/methods-read-proxy", map[string]string{
		"address_hash": addressHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) GetSmartContractWriteMethodsProxy(ctx context.Context, addressHash string) (resp []SmartContractMethod, err error) {
	path := buildPath("/api/v2/smart-contracts/{address_hash}/methods-write-proxy", map[string]string{
		"address_hash": addressHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)
	return
}

func (c *client) QuerySmartContractReadMethod(ctx context.Context, addressHash string, reqBody QueryReadMethodReq) (resp []SmartContractReadResult, err error) {
	path := buildPath("/api/v2/smart-contracts/{address_hash}/query-read-method", map[string]string{
		"address_hash": addressHash,
	})
	_, err = c.resty.R().
		SetContext(ctx).
		SetBody(reqBody).
		SetResult(&resp).
		Post(path)
	return
}
