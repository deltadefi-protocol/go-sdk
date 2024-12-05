package client

import (
	"encoding/json"
	"fmt"

	"github.com/deltadefi-protocol/go-sdk/pkg/api/requests"
	"github.com/deltadefi-protocol/go-sdk/pkg/api/responses"
)

func (c *Client) getDepth(data *requests.GetMarketDepthRequest) (*responses.GetMarketDepthResponse, error) {
	url := fmt.Sprintf("/market/depth?pair=%s", data.Pair)
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getMarketDepthResponse responses.GetMarketDepthResponse
	err = json.NewDecoder(resp.Body).Decode(&getMarketDepthResponse)
	if err != nil {
		return nil, err
	}

	return &getMarketDepthResponse, nil
}

func (c *Client) getMarketPrice(data *requests.GetMarketPriceRequest) (*responses.GetMarketPriceResponse, error) {
	url := fmt.Sprintf("/market/market-price?pair=%s", data.Pair)
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getMarketPriceResponse responses.GetMarketPriceResponse
	err = json.NewDecoder(resp.Body).Decode(&getMarketPriceResponse)
	if err != nil {
		return nil, err
	}

	return &getMarketPriceResponse, nil
}

func (c *Client) getAggregatedPrice(data *requests.GetAggregatedPriceRequest) (*responses.GetAggregatedPriceResponse, error) {
	var url string
	if *data.Start == 0 && *data.End == 0 {
		url = fmt.Sprintf("/market/aggregate/%s?interval=%s", data.Pair, data.Interval)
	} else {
		url = fmt.Sprintf("/market/aggregate/%s?interval=%s&start=%s&end=%s", data.Pair, data.Interval, data.Start, data.End)
	}
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getAggregatedPriceResponse responses.GetAggregatedPriceResponse
	err = json.NewDecoder(resp.Body).Decode(&getAggregatedPriceResponse)
	if err != nil {
		return nil, err
	}

	return &getAggregatedPriceResponse, nil
}
