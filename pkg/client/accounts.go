package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/deltadefi-protocol/go-sdk/pkg/api/requests"
	"github.com/deltadefi-protocol/go-sdk/pkg/api/responses"
)

func (c *Client) SignIn(data *requests.SignInRequest) (*responses.SignInResponse, error) {
	requestBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/accounts/signin", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", c.apiKey)
	req.Header.Set("Authorization", c.Jwt)
	req.Header.Set("auth_key", data.AuthKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var signInResponse responses.SignInResponse
	err = json.NewDecoder(resp.Body).Decode(&signInResponse)
	if err != nil {
		return nil, err
	}

	return &signInResponse, nil
}

func (c *Client) GetDepositRecords() (*responses.GetDepositRecordsResponse, error) {
	resp, err := c.get("/accounts/deposit-records")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getDepositRecordsResponse responses.GetDepositRecordsResponse
	err = json.NewDecoder(resp.Body).Decode(&getDepositRecordsResponse)
	if err != nil {
		return nil, err
	}

	return &getDepositRecordsResponse, nil
}

func (c *Client) GetWithdrawalRecords() (*responses.GetWithdrawalRecordsResponse, error) {
	resp, err := c.get("/accounts/withdrawal-records")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getWithdrawalRecordsResponse responses.GetWithdrawalRecordsResponse
	err = json.NewDecoder(resp.Body).Decode(&getWithdrawalRecordsResponse)
	if err != nil {
		return nil, err
	}

	return &getWithdrawalRecordsResponse, nil
}

func (c *Client) GetOrderRecords() (*responses.GetOrderRecordResponse, error) {
	resp, err := c.get("/accounts/orde-records")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getOrderRecordResponse responses.GetOrderRecordResponse
	err = json.NewDecoder(resp.Body).Decode(&getOrderRecordResponse)
	if err != nil {
		return nil, err
	}

	return &getOrderRecordResponse, nil
}

func (c *Client) GetAccountBalance() (*responses.GetAccountBalanceResponse, error) {
	resp, err := c.get("/accounts/balance")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getAccountBalanceResponse responses.GetAccountBalanceResponse
	err = json.NewDecoder(resp.Body).Decode(&getAccountBalanceResponse)
	if err != nil {
		return nil, err
	}

	return &getAccountBalanceResponse, nil
}

func (c *Client) CreateNewAPIKey() (*responses.GenerateNewAPIKeyResponse, error) {
	resp, err := c.get("/accounts/new-api-key")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var generateNewAPIKeyResponse responses.GenerateNewAPIKeyResponse
	err = json.NewDecoder(resp.Body).Decode(&generateNewAPIKeyResponse)
	if err != nil {
		return nil, err
	}

	return &generateNewAPIKeyResponse, nil
}

func (c *Client) BuildDepositTransaction(data *requests.BuildDepositTransactionRequest) (*responses.BuildDepositTransactionResponse, error) {
	requestBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/accounts/deposit/build", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", c.apiKey)
	req.Header.Set("Authorization", c.Jwt)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buildDepositTransactionResponse responses.BuildDepositTransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&buildDepositTransactionResponse)
	if err != nil {
		return nil, err
	}

	return &buildDepositTransactionResponse, nil
}

func (c *Client) BuildWithdrawalTransaction(data *requests.BuildWithdrawalTransactionRequest) (*responses.BuildWithdrawalTransactionResponse, error) {
	resp, err := c.post("/accounts/withdrawal/build", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buildWithdrawalTransactionResponse responses.BuildWithdrawalTransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&buildWithdrawalTransactionResponse)
	if err != nil {
		return nil, err
	}

	return &buildWithdrawalTransactionResponse, nil
}

func (c *Client) SubmitDepositTransaction(data *requests.SubmitDepositTransactionRequest) (*responses.SubmitDepositTransactionResponse, error) {
	resp, err := c.post("/accounts/deposit/submit", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var submitDepositTransactionResponse responses.SubmitDepositTransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&submitDepositTransactionResponse)
	if err != nil {
		return nil, err
	}

	return &submitDepositTransactionResponse, nil
}

func (c *Client) SubmitWithdrawalTransaction(data *requests.SubmitWithdrawalTransactionRequest) (*responses.SubmitWithdrawalTransactionResponse, error) {
	resp, err := c.post("/accounts/withdrawal/submit", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var submitWithdrawalTransactionResponse responses.SubmitWithdrawalTransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&submitWithdrawalTransactionResponse)
	if err != nil {
		return nil, err
	}

	return &submitWithdrawalTransactionResponse, nil
}

func (c *Client) GetTermsAndCondition() (*responses.GetTermsAndConditionResponse, error) {
	resp, err := c.get("/accounts/terms-and-condition")
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
