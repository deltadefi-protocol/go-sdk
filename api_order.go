package deltadefi

// func (c *Client) buildPlaceOrderTransaction(data *BuildPlaceOrderTransactionRequest) (*BuildPlaceOrderTransactionResponse, error) {
// 	resp, err := c.post("/order/build", data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var buildPlaceOrderTransactionResponse BuildPlaceOrderTransactionResponse
// 	err = json.NewDecoder(resp.Body).Decode(&buildPlaceOrderTransactionResponse)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &buildPlaceOrderTransactionResponse, nil
// }

// func (c *Client) buildCancelOrderTransaction(orderId string) (*BuildCancelOrderTransactionResponse, error) {
// 	url := fmt.Sprintf("/order/%s/build", orderId)
// 	var empty interface{}
// 	resp, err := c.delete(url, empty)

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var buildCancelOrderTransactionResponse BuildCancelOrderTransactionResponse
// 	err = json.NewDecoder(resp.Body).Decode(&buildCancelOrderTransactionResponse)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &buildCancelOrderTransactionResponse, nil
// }

// func (c *Client) submitPlaceOrderTransactionRequest(data *SubmitPlaceOrderTransactionRequest) (*SubmitPlaceOrderTransactionResponse, error) {
// 	resp, err := c.post("/order/submit", data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var submitPlaceOrderTransactionResponse SubmitPlaceOrderTransactionResponse
// 	err = json.NewDecoder(resp.Body).Decode(&submitPlaceOrderTransactionResponse)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &submitPlaceOrderTransactionResponse, nil
// }

// func (c *Client) submitCancelOrderTransactionRequest(data *SubmitCancelOrderTransactionRequest) (*SubmitCancelOrderTransactionResponse, error) {
// 	resp, err := c.delete("/order/submit", data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var submitCancelOrderTransactionResponse SubmitCancelOrderTransactionResponse
// 	err = json.NewDecoder(resp.Body).Decode(&submitCancelOrderTransactionResponse)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &submitCancelOrderTransactionResponse, nil
// }
