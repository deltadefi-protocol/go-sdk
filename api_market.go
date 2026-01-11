package deltadefi

import (
	"encoding/json"
	"fmt"
)

// MarketClient provides access to market data operations.
type MarketClient struct {
	pathUrl string
	client  *Client
}

// newMarketClient creates a new MarketClient instance.
func newMarketClient(client *Client) *MarketClient {
	return &MarketClient{
		pathUrl: "/market",
		client:  client,
	}
}

// GetMarketPrice retrieves the current market price for the specified trading pair.
//
// Parameters:
//   - symbol: Trading pair symbol (e.g., "ADAUSDM")
//
// Returns:
//   - *GetMarketPriceResponse: Current market price
//   - error: nil on success, error on failure
func (c *MarketClient) GetMarketPrice(symbol string) (*GetMarketPriceResponse, error) {
	params := make(map[string]string)
	params["symbol"] = symbol

	bodyBytes, err := c.client.getWithParams(c.pathUrl+"/market-price", params)
	if err != nil {
		return nil, err
	}

	var getMarketPriceResponse GetMarketPriceResponse
	err = json.Unmarshal(bodyBytes, &getMarketPriceResponse)
	if err != nil {
		return nil, err
	}
	return &getMarketPriceResponse, nil
}

// GetAggregatedPrice retrieves historical price data (candlesticks) for the specified parameters.
// Supports various time intervals and date ranges for technical analysis.
//
// Parameters:
//   - data: Request parameters including symbol, interval, start and end timestamps
//
// Returns:
//   - *GetAggregatedPriceResponse: Array of candlestick data (OHLCV)
//   - error: nil on success, error on failure
func (c *MarketClient) GetAggregatedPrice(data *GetAggregatedPriceRequest) (*GetAggregatedPriceResponse, error) {
	fullPath := c.pathUrl + "/aggregate/" + string(data.Symbol) + "?interval=" + string(data.Interval) +
		"&start=" + fmt.Sprint(data.Start) + "&end=" + fmt.Sprint(data.End)
	bodyBytes, err := c.client.get(fullPath)
	if err != nil {
		return nil, err
	}

	var getAggregatedPriceResponse GetAggregatedPriceResponse
	err = json.Unmarshal(bodyBytes, &getAggregatedPriceResponse)
	if err != nil {
		return nil, err
	}
	return &getAggregatedPriceResponse, nil
}

// GetTickerPrice retrieves the ticker price for the specified trading pair.
//
// Parameters:
//   - symbol: Trading pair symbol (e.g., "ADAUSDM")
//
// Returns:
//   - *GetTickerPriceResponse: Current ticker price
//   - error: nil on success, error on failure
func (c *MarketClient) GetTickerPrice(symbol string) (*GetTickerPriceResponse, error) {
	bodyBytes, err := c.client.get(c.pathUrl + "/ticker-price/" + symbol)
	if err != nil {
		return nil, err
	}

	var getTickerPriceResponse GetTickerPriceResponse
	err = json.Unmarshal(bodyBytes, &getTickerPriceResponse)
	if err != nil {
		return nil, err
	}
	return &getTickerPriceResponse, nil
}

// GetOrderbookDepth retrieves the orderbook depth for the specified trading pair.
//
// Parameters:
//   - symbol: Trading pair symbol (e.g., "ADAUSDM")
//
// Returns:
//   - *GetMarketDepthResponse: Current orderbook with bids and asks
//   - error: nil on success, error on failure
func (c *MarketClient) GetOrderbookDepth(symbol string) (*GetMarketDepthResponse, error) {
	bodyBytes, err := c.client.get(c.pathUrl + "/orderbook-depth/" + symbol)
	if err != nil {
		return nil, err
	}

	var getMarketDepthResponse GetMarketDepthResponse
	err = json.Unmarshal(bodyBytes, &getMarketDepthResponse)
	if err != nil {
		return nil, err
	}
	return &getMarketDepthResponse, nil
}
