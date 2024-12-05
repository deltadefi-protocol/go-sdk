package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/deltadefi-protocol/go-sdk/pkg/config"
)

type Client struct {
	apiKey     string
	NetworkId  uint8
	Jwt        string
	SigningKey string
	HTTPClient *http.Client
	BaseUrl    string
}

func NewClient(apiKey string, network string, jwt string, signingKey string) *Client {
	cfg := config.GetConfig()
	var networkId uint8
	var baseUrl string

	if network == "mainnet" {
		networkId = uint8(1)
		baseUrl = "https://api-dev.deltadefi.io" // TODO: input production link once available
	} else {
		networkId = uint8(0)
		baseUrl = "https://api-dev.deltadefi.io"
	}

	return &Client{
		apiKey:     apiKey,
		NetworkId:  networkId,
		Jwt:        jwt,
		SigningKey: signingKey,
		HTTPClient: &http.Client{
			Timeout: time.Duration(cfg.Client.Timeout) * time.Minute,
		},
		BaseUrl: baseUrl,
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

	req.Header.Set("Accept", "application/json")
	req.Header.Add("api-key", c.apiKey)
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
	req, err := http.NewRequest("GET", c.BaseUrl+url, nil)
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Add("api-key", c.apiKey)
	req.Header.Set("Authorization", c.Jwt)

	return c.HTTPClient.Do(req)
}

func (c *Client) post(url string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.BaseUrl+url, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Add("api-key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.Jwt)

	return c.HTTPClient.Do(req)
}

func (c *Client) postBuffer(url string, buffer []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", c.BaseUrl+url, bytes.NewBuffer(buffer))
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}

	req.Header.Set("Accept", "application/cbor")
	req.Header.Add("api-key", c.apiKey)
	req.Header.Set("Content-Type", "application/cbor")
	req.Header.Set("Authorization", c.Jwt)

	return c.HTTPClient.Do(req)
}
