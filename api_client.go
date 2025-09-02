// Package deltadefi provides a Go SDK for interacting with the DeltaDeFi API.
// DeltaDeFi is a decentralized finance protocol built on Cardano that enables
// trading operations, account management, and market data access.
package deltadefi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	rum "github.com/sidan-lab/rum/wallet"
)

// DeltaDeFi is the main client for interacting with the DeltaDeFi API.
// It provides access to all API endpoints through specialized client instances.
type DeltaDeFi struct {
	// Accounts provides access to account management operations
	Accounts        *AccountsClient
	// Market provides access to market data operations
	Market          *MarketClient
	// Order provides access to order management operations
	Order           *OrderClient
	// MasterWallet holds the master wallet instance
	MasterWallet    *rum.Wallet
	// OperationWallet holds the operation wallet instance for transaction signing
	OperationWallet *rum.Wallet
	// client is the underlying HTTP client
	client          *Client
}

// NewDeltaDeFi creates a new DeltaDeFi client instance.
// It initializes the HTTP client and all API endpoint clients.
//
// Parameters:
//   - cfg: ApiConfig containing network, API key, and operation passcode
//
// Returns:
//   - *DeltaDeFi: A new client instance ready for API operations
func NewDeltaDeFi(cfg ApiConfig) *DeltaDeFi {
	client := newClient(cfg)
	return &DeltaDeFi{
		Accounts:        newAccountsClient(client),
		Market:          newMarketClient(client),
		Order:           newOrderClient(client),
		MasterWallet:    nil,
		OperationWallet: nil,
		client:          client,
	}
}

// Client represents the underlying HTTP client for API communication.
type Client struct {
	// ApiKey is the API authentication key
	ApiKey            string
	// NetworkId identifies the network (0=dev/staging, 1=mainnet)
	NetworkId         uint8
	// OperationPasscode is used for decrypting operation keys
	OperationPasscode string
	// HTTPClient is the HTTP client instance
	HTTPClient        *http.Client
	// BaseURL is the API base URL
	BaseURL           string
}

// newClient creates a new HTTP client instance based on the provided configuration.
// It sets the appropriate base URL and network ID based on the network selection.
func newClient(cfg ApiConfig) *Client {
	var networkId uint8
	var baseURL string

	if cfg.Network == "mainnet" {
		networkId = uint8(1)
		baseURL = "https://api-staging.deltadefi.io" // TODO: input production link once available
	} else if cfg.Network == "staging" {
		networkId = uint8(0)
		baseURL = "https://api-staging.deltadefi.io"
	} else {
		networkId = uint8(0)
		baseURL = "https://api-dev.deltadefi.io"
	}

	if (cfg.ProvidedBaseUrl) != "" {
		baseURL = cfg.ProvidedBaseUrl
	}

	return &Client{
		ApiKey:            cfg.ApiKey,
		NetworkId:         networkId,
		OperationPasscode: cfg.OperationPasscode,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		BaseURL: baseURL,
	}
}

// get performs a GET request to the specified URL path.
// It automatically adds authentication headers and returns the response body.
func (c *Client) get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.BaseURL+url, nil)
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", c.ApiKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

// getWithParams performs a GET request with query parameters.
// It automatically adds authentication headers and handles parameter encoding.
func (c *Client) getWithParams(path string, params map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.BaseURL+path, nil)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", c.ApiKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check if the response status code is not 2xx
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API error: %s, status code: %d", string(bodyBytes), resp.StatusCode)
	}

	return bodyBytes, nil
}

// post performs a POST request with JSON body.
// It automatically adds authentication headers and marshals the request body.
func (c *Client) post(url string, body interface{}) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.BaseURL+url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", c.ApiKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

// delete performs a DELETE request with JSON body.
// It automatically adds authentication headers and marshals the request body.
func (c *Client) delete(url string, body interface{}) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("DELETE", c.BaseURL+url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", c.ApiKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
