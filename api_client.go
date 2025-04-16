package deltadefi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type DeltaDeFi struct {
	Account *AccountClient
	App     *AppClient
	Market  *MarketClient
}

func NewDeltaDeFi(cfg ApiConfig) *DeltaDeFi {
	client := newClient(cfg)
	return &DeltaDeFi{
		Account: newAccountClient(client),
		App:     newAppClient(client),
		Market:  newMarketClient(client),
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

func (c *Client) delete(url string, body interface{}) (*http.Response, error) {
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

	return c.HTTPClient.Do(req)
}
