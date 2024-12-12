package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/deltadefi-protocol/go-sdk/pkg/api"
)

type Client struct {
	apiKey     string
	NetworkId  uint8
	Jwt        string
	SigningKey string
	HTTPClient *http.Client
	BaseURL    string
}

func NewClient(cfg api.ApiConfig, ProvidedBaseURL string) *Client {
	var networkId uint8
	var baseURL string

	if *cfg.Network == "mainnet" {
		networkId = uint8(1)
		baseURL = "https://api-dev.deltadefi.io" // TODO: input production link once available
	} else if *cfg.Network == "preprod" {
		networkId = uint8(0)
		baseURL = "https://api-dev.deltadefi.io"
	} else {
		panic("unsupported network")
	}

	if (ProvidedBaseURL) != "" {
		baseURL = ProvidedBaseURL
	}

	return &Client{
		apiKey:     *cfg.APIKey,
		NetworkId:  networkId,
		Jwt:        *cfg.JWT,
		SigningKey: *cfg.SigningKey,
		HTTPClient: &http.Client{},
		BaseURL:    baseURL,
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// type successResponse struct {
// 	Code int         `json:"code"`
// 	Data interface{} `json:"data"`
// }

func (c *Client) sendRequest(req *http.Request, responseBody *string) error {
	if req == nil {
		return fmt.Errorf("empty request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", c.apiKey)
	req.Header.Set("Authorization", c.Jwt)

	if c.HTTPClient == nil {
		return fmt.Errorf("missing http client")
	}
	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return err
	}
	if resp == nil {
		return fmt.Errorf("empty response")
	}

	// Try to unmarshall into errorResponse
	if resp.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(resp.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", resp.StatusCode)
	}

	respBodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("failed to read body: %s", err)
	}
	defer resp.Body.Close()

	*responseBody = string(respBodyBytes)

	return nil
}

func (c *Client) get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+url, nil)
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", c.apiKey)
	req.Header.Set("Authorization", c.Jwt)

	return c.HTTPClient.Do(req)
}

func (c *Client) post(url string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.BaseURL+url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", c.apiKey)
	req.Header.Set("Authorization", c.Jwt)

	return c.HTTPClient.Do(req)
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
	req.Header.Add("X-API-KEY", c.apiKey)
	req.Header.Set("Authorization", c.Jwt)

	return c.HTTPClient.Do(req)
}

// func (c *Client) postBuffer(url string, buffer []byte) (*http.Response, error) {
// 	req, err := http.NewRequest("POST", c.BaseURL+url, bytes.NewBuffer(buffer))
// 	if err != nil {
// 		return nil, err
// 	}
// 	if req == nil {
// 		return nil, fmt.Errorf("empty request")
// 	}

// 	req.Header.Set("Accept", "application/cbor")
// 	req.Header.Add("X-API-KEY", c.apiKey)
// 	req.Header.Set("Content-Type", "application/cbor")
// 	req.Header.Set("Authorization", c.Jwt)

// 	return c.HTTPClient.Do(req)
// }
