package client

import (
	"encoding/json"
	"fmt"

	"github.com/deltadefi-protocol/go-sdk/pkg/api/requests"
	"github.com/deltadefi-protocol/go-sdk/pkg/api/responses"
)

func (c *Client) buildPlaceOrderTransaction(data *requests.BuildPlaceOrderTransactionRequest) (*responses.BuildPlaceOrderTransactionResponse, error) {
	resp, err := c.post("/order/build", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buildPlaceOrderTransactionResponse responses.BuildPlaceOrderTransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&buildPlaceOrderTransactionResponse)
	if err != nil {
		return nil, err
	}

	return &buildPlaceOrderTransactionResponse, nil
}

func (c *Client) buildCancelOrderTransaction(orderId string) (*responses.BuildCancelOrderTransactionResponse, error) {
	url := fmt.Sprintf("/order/%s/build", orderId)
	var empty interface{}
	resp, err := c.delete(url, empty)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buildCancelOrderTransactionResponse responses.BuildCancelOrderTransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&buildCancelOrderTransactionResponse)
	if err != nil {
		return nil, err
	}

	return &buildCancelOrderTransactionResponse, nil
}

func (c *Client) submitPlaceOrderTransactionRequest(data *requests.SubmitPlaceOrderTransactionRequest) (*responses.SubmitPlaceOrderTransactionResponse, error) {
	resp, err := c.post("/order/submit", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var submitPlaceOrderTransactionResponse responses.SubmitPlaceOrderTransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&submitPlaceOrderTransactionResponse)
	if err != nil {
		return nil, err
	}

	return &submitPlaceOrderTransactionResponse, nil
}

func (c *Client) submitCancelOrderTransactionRequest(data *requests.SubmitCancelOrderTransactionRequest) (*responses.SubmitCancelOrderTransactionResponse, error) {
	resp, err := c.delete("/order/submit", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var submitCancelOrderTransactionResponse responses.SubmitCancelOrderTransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&submitCancelOrderTransactionResponse)
	if err != nil {
		return nil, err
	}

	return &submitCancelOrderTransactionResponse, nil
}
