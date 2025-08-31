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

type DeltaDeFi struct {
	Accounts        *AccountsClient
	Market          *MarketClient
	Order           *OrderClient
	MasterWallet    *rum.Wallet
	OperationWallet *rum.Wallet
	client          *Client
}

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

type Client struct {
	ApiKey            string
	NetworkId         uint8
	OperationPasscode string
	HTTPClient        *http.Client
	BaseURL           string
}

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
