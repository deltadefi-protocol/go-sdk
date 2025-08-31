package deltadefi

import (
	"encoding/json"
)

type OrderClient struct {
	pathUrl string
	client  *Client
}

func newOrderClient(client *Client) *OrderClient {
	return &OrderClient{
		pathUrl: "/order",
		client:  client,
	}
}

func (c *OrderClient) BuildPlaceOrderTransaction(data *BuildPlaceOrderTransactionRequest) (*BuildPlaceOrderTransactionResponse, error) {
	bodyBytes, err := c.client.post(c.pathUrl+"/build", data)
	if err != nil {
		return nil, err
	}

	var buildPlaceOrderTransactionResponse BuildPlaceOrderTransactionResponse
	err = json.Unmarshal(bodyBytes, &buildPlaceOrderTransactionResponse)
	if err != nil {
		return nil, err
	}
	return &buildPlaceOrderTransactionResponse, nil
}

func (c *OrderClient) BuildCancelOrderTransaction(orderId string) (*BuildCancelOrderTransactionResponse, error) {
	bodyBytes, err := c.client.delete(c.pathUrl+"/"+orderId+"/build", nil)
	if err != nil {
		return nil, err
	}

	var buildCancelOrderTransactionResponse BuildCancelOrderTransactionResponse
	err = json.Unmarshal(bodyBytes, &buildCancelOrderTransactionResponse)
	if err != nil {
		return nil, err
	}
	return &buildCancelOrderTransactionResponse, nil
}

func (c *OrderClient) SubmitPlaceOrderTransactionRequest(data *SubmitPlaceOrderTransactionRequest) (*SubmitPlaceOrderTransactionResponse, error) {
	bodyBytes, err := c.client.post(c.pathUrl+"/submit", data)
	if err != nil {
		return nil, err
	}

	var submitPlaceOrderTransactionResponse SubmitPlaceOrderTransactionResponse
	err = json.Unmarshal(bodyBytes, &submitPlaceOrderTransactionResponse)
	if err != nil {
		return nil, err
	}
	return &submitPlaceOrderTransactionResponse, nil
}

func (c *OrderClient) SubmitCancelOrderTransactionRequest(data *SubmitCancelOrderTransactionRequest) (*SubmitCancelOrderTransactionResponse, error) {
	bodyBytes, err := c.client.delete(c.pathUrl+"/submit", data)
	if err != nil {
		return nil, err
	}

	var submitCancelOrderTransactionResponse SubmitCancelOrderTransactionResponse
	err = json.Unmarshal(bodyBytes, &submitCancelOrderTransactionResponse)
	if err != nil {
		return nil, err
	}
	return &submitCancelOrderTransactionResponse, nil
}
