package deltadefi

// func (c *Client) getDepth(data *GetMarketDepthRequest) (*GetMarketDepthResponse, error) {
// 	url := fmt.Sprintf("/market/depth?pair=%s", data.Pair)
// 	resp, err := c.get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var getMarketDepthResponse GetMarketDepthResponse
// 	err = json.NewDecoder(resp.Body).Decode(&getMarketDepthResponse)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &getMarketDepthResponse, nil
// }

// func (c *Client) getMarketPrice(data *GetMarketPriceRequest) (*GetMarketPriceResponse, error) {
// 	url := fmt.Sprintf("/market/market-price?pair=%s", data.Pair)
// 	resp, err := c.get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var getMarketPriceResponse GetMarketPriceResponse
// 	err = json.NewDecoder(resp.Body).Decode(&getMarketPriceResponse)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &getMarketPriceResponse, nil
// }

// func (c *Client) getAggregatedPrice(data *GetAggregatedPriceRequest) (*GetAggregatedPriceResponse, error) {
// 	var url string
// 	if *data.Start == 0 && *data.End == 0 {
// 		url = fmt.Sprintf("/market/aggregate/%s?interval=%s", data.Pair, data.Interval)
// 	} else {
// 		url = fmt.Sprintf("/market/aggregate/%s?interval=%s&start=%d&end=%d", data.Pair, data.Interval, *data.Start, *data.End)
// 	}
// 	resp, err := c.get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var getAggregatedPriceResponse GetAggregatedPriceResponse
// 	err = json.NewDecoder(resp.Body).Decode(&getAggregatedPriceResponse)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &getAggregatedPriceResponse, nil
// }
