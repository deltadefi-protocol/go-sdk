package deltadefi

import (
	"encoding/json"
)

// OrderClient provides access to order management operations.
type OrderClient struct {
	pathUrl string
	client  *Client
}

// newOrderClient creates a new OrderClient instance.
func newOrderClient(client *Client) *OrderClient {
	return &OrderClient{
		pathUrl: "/order",
		client:  client,
	}
}

// BuildPlaceOrderTransaction builds a transaction for placing a new order.
// Supports both market and limit orders with optional slippage controls.
// The returned transaction hex must be signed and then submitted using SubmitPlaceOrderTransactionRequest.
//
// Parameters:
//   - data: Order details including symbol, side, type, quantity, and optional price/slippage settings
//
// Returns:
//   - *BuildPlaceOrderTransactionResponse: Order ID and transaction hex ready for signing
//   - error: nil on success, error on failure
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

// BuildCancelOrderTransaction builds a transaction for canceling an existing order.
// The returned transaction hex must be signed and then submitted using SubmitCancelOrderTransactionRequest.
//
// Parameters:
//   - orderId: The unique identifier of the order to cancel
//
// Returns:
//   - *BuildCancelOrderTransactionResponse: Transaction hex ready for signing
//   - error: nil on success, error on failure
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

// SubmitPlaceOrderTransactionRequest submits a signed place order transaction to the network.
//
// Parameters:
//   - data: Submit request containing the order ID and signed transaction hex
//
// Returns:
//   - *SubmitPlaceOrderTransactionResponse: Complete order details after submission
//   - error: nil on success, error on failure
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

// SubmitCancelOrderTransactionRequest submits a signed cancel order transaction to the network.
//
// Parameters:
//   - data: Submit request containing the signed transaction hex
//
// Returns:
//   - *SubmitCancelOrderTransactionResponse: Transaction hash of the cancellation
//   - error: nil on success, error on failure
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
