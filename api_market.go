package deltadefi

import (
	"encoding/json"
	"fmt"
)

type MarketClient struct {
	pathUrl string
	client  *Client
}

func newMarketClient(client *Client) *MarketClient {
	return &MarketClient{
		pathUrl: "/market",
		client:  client,
	}
}

func (c *MarketClient) GetDepth(symbol string) (*GetMarketDepthResponse, error) {
	bodyBytes, err := c.client.get(c.pathUrl + "/depth?symbol=" + symbol)
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

func (c *MarketClient) GetMarketPrice(symbol string) (*GetMarketPriceResponse, error) {
	bodyBytes, err := c.client.get(c.pathUrl + "/market-price?symbol=" + symbol)
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

func (c *MarketClient) GetAggregatedPrice(data *GetAggregatedPriceRequest) (*GetAggregatedPriceResponse, error) {
	bodyBytes, err := c.client.get(
		c.pathUrl + "/aggregated-trade/" + data.Symbol + "?interval=" + string(data.Interval) +
			"&start=" + fmt.Sprint(data.Start) + "&end=" + fmt.Sprint(data.End))
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
