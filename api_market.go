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

func (c *MarketClient) GetAggregatedPrice(data *GetAggregatedPriceRequest) (*GetAggregatedPriceResponse, error) {
	fullPath := c.pathUrl + "/graph/" + string(data.Symbol) + "?interval=" + string(data.Interval) +
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
