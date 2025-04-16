package deltadefi

import (
	"encoding/json"
)

type AccountsClient struct {
	pathUrl string
	client  *Client
}

func newAccountsClient(client *Client) *AccountsClient {
	return &AccountsClient{
		pathUrl: "/accounts",
		client:  client,
	}
}

func (c *AccountsClient) GetOperationKey() (*GetOperationKeyResponse, error) {
	bodyBytes, err := c.client.get(c.pathUrl + "/operation-key")
	if err != nil {
		return nil, err
	}

	var getOperationKeyResponse GetOperationKeyResponse
	err = json.Unmarshal(bodyBytes, &getOperationKeyResponse)
	if err != nil {
		return nil, err
	}
	return &getOperationKeyResponse, nil
}

func (c *AccountsClient) CreateNewAPIKey() (*CreateNewAPIKeyResponse, error) {
	bodyBytes, err := c.client.get(c.pathUrl + "/new-api-key")
	if err != nil {
		return nil, err
	}

	var createNewAPIKeyResponse CreateNewAPIKeyResponse
	err = json.Unmarshal(bodyBytes, &createNewAPIKeyResponse)
	if err != nil {
		return nil, err
	}
	return &createNewAPIKeyResponse, nil
}

func (c *AccountsClient) GetDepositRecords() (*GetDepositRecordsResponse, error) {
	bodyBytes, err := c.client.get(c.pathUrl + "/deposit-records")
	if err != nil {
		return nil, err
	}

	var getDepositRecordsResponse GetDepositRecordsResponse
	err = json.Unmarshal(bodyBytes, &getDepositRecordsResponse)
	if err != nil {
		return nil, err
	}
	return &getDepositRecordsResponse, nil
}

func (c *AccountsClient) GetWithdrawalRecords() (*GetWithdrawalRecordsResponse, error) {
	bodyBytes, err := c.client.get(c.pathUrl + "/withdrawal-records")
	if err != nil {
		return nil, err
	}

	var getWithdrawalRecordsResponse GetWithdrawalRecordsResponse
	err = json.Unmarshal(bodyBytes, &getWithdrawalRecordsResponse)
	if err != nil {
		return nil, err
	}
	return &getWithdrawalRecordsResponse, nil
}

func (c *AccountsClient) GetOrderRecords() (*GetOrderRecordResponse, error) {
	bodyBytes, err := c.client.get(c.pathUrl + "/order-records")
	if err != nil {
		return nil, err
	}

	var getOrderRecordResponse GetOrderRecordResponse
	err = json.Unmarshal(bodyBytes, &getOrderRecordResponse)
	if err != nil {
		return nil, err
	}
	return &getOrderRecordResponse, nil
}

func (c *AccountsClient) GetAccountBalance() (*GetAccountBalanceResponse, error) {
	bodyBytes, err := c.client.get(c.pathUrl + "/balance")
	if err != nil {
		return nil, err
	}

	var getAccountBalanceResponse GetAccountBalanceResponse
	err = json.Unmarshal(bodyBytes, &getAccountBalanceResponse)
	if err != nil {
		return nil, err
	}
	return &getAccountBalanceResponse, nil
}

func (c *AccountsClient) BuildDepositTransaction(data *BuildDepositTransactionRequest) (*BuildDepositTransactionResponse, error) {
	bodyBytes, err := c.client.post(c.pathUrl+"/deposit/build", data)
	if err != nil {
		return nil, err
	}

	var buildDepositTransactionResponse BuildDepositTransactionResponse
	err = json.Unmarshal(bodyBytes, &buildDepositTransactionResponse)
	if err != nil {
		return nil, err
	}
	return &buildDepositTransactionResponse, nil
}

func (c *AccountsClient) BuildWithdrawalTransaction(data *BuildWithdrawalTransactionRequest) (*BuildWithdrawalTransactionResponse, error) {
	bodyBytes, err := c.client.post(c.pathUrl+"/withdrawal/build", data)
	if err != nil {
		return nil, err
	}

	var buildWithdrawalTransactionResponse BuildWithdrawalTransactionResponse
	err = json.Unmarshal(bodyBytes, &buildWithdrawalTransactionResponse)
	if err != nil {
		return nil, err
	}
	return &buildWithdrawalTransactionResponse, nil
}

func (c *AccountsClient) SubmitDepositTransaction(data *SubmitDepositTransactionRequest) (*SubmitDepositTransactionResponse, error) {
	bodyBytes, err := c.client.post(c.pathUrl+"/deposit/submit", data)
	if err != nil {
		return nil, err
	}

	var submitDepositTransactionResponse SubmitDepositTransactionResponse
	err = json.Unmarshal(bodyBytes, &submitDepositTransactionResponse)
	if err != nil {
		return nil, err
	}
	return &submitDepositTransactionResponse, nil
}

func (c *AccountsClient) SubmitWithdrawalTransaction(data *SubmitWithdrawalTransactionRequest) (*SubmitWithdrawalTransactionResponse, error) {
	bodyBytes, err := c.client.post(c.pathUrl+"/withdrawal/submit", data)
	if err != nil {
		return nil, err
	}

	var submitWithdrawalTransactionResponse SubmitWithdrawalTransactionResponse
	err = json.Unmarshal(bodyBytes, &submitWithdrawalTransactionResponse)
	if err != nil {
		return nil, err
	}
	return &submitWithdrawalTransactionResponse, nil
}
