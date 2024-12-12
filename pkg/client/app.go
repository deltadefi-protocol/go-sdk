package client

import (
	"encoding/json"

	"github.com/deltadefi-protocol/go-sdk/pkg/api/responses"
)

func (c *Client) getTermsAndCondition() (*responses.GetTermsAndConditionResponse, error) {
	resp, err := c.get("/terms-and-conditions")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getTermsAndConditionResponse responses.GetTermsAndConditionResponse
	err = json.NewDecoder(resp.Body).Decode(&getTermsAndConditionResponse)
	if err != nil {
		return nil, err
	}

	return &getTermsAndConditionResponse, nil
}
