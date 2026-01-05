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
//   - data: Order details including symbol, side, type, quantity, and optional price/slippage/post-only settings
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

// CancelOrder cancels an existing order by order ID.
// This is a simplified endpoint that handles cancellation directly without build/sign/submit flow.
//
// Parameters:
//   - orderId: The unique identifier of the order to cancel
//
// Returns:
//   - *CancelOrderResponse: Contains the cancelled order ID
//   - error: nil on success, error on failure
func (c *OrderClient) CancelOrder(orderId string) (*CancelOrderResponse, error) {
	bodyBytes, err := c.client.post(c.pathUrl+"/"+orderId+"/cancel", nil)
	if err != nil {
		return nil, err
	}

	var cancelOrderResponse CancelOrderResponse
	err = json.Unmarshal(bodyBytes, &cancelOrderResponse)
	if err != nil {
		return nil, err
	}
	return &cancelOrderResponse, nil
}

// CancelAllOrders cancels all open orders for a given symbol.
// This is a simplified endpoint that handles cancellation directly without build/sign/submit flow.
//
// Parameters:
//   - symbol: The trading pair symbol to cancel orders for
//
// Returns:
//   - *CancelAllOrdersResponse: Contains the symbol and list of cancelled order IDs
//   - error: nil on success, error on failure
func (c *OrderClient) CancelAllOrders(symbol string) (*CancelAllOrdersResponse, error) {
	bodyBytes, err := c.client.post(c.pathUrl+"/cancel-all", &CancelAllOrdersRequest{Symbol: symbol})
	if err != nil {
		return nil, err
	}

	var cancelAllOrdersResponse CancelAllOrdersResponse
	err = json.Unmarshal(bodyBytes, &cancelAllOrdersResponse)
	if err != nil {
		return nil, err
	}
	return &cancelAllOrdersResponse, nil
}

// SubmitPlaceOrderTransaction submits a signed place order transaction to the network.
//
// Parameters:
//   - data: Submit request containing the order ID and signed transaction hex
//
// Returns:
//   - *SubmitPlaceOrderTransactionResponse: Complete order details after submission
//   - error: nil on success, error on failure
func (c *OrderClient) SubmitPlaceOrderTransaction(data *SubmitPlaceOrderTransactionRequest) (*SubmitPlaceOrderTransactionResponse, error) {
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

