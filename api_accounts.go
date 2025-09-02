package deltadefi

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// AccountsClient provides access to account management operations.
type AccountsClient struct {
	pathUrl string
	client  *Client
}

// newAccountsClient creates a new AccountsClient instance.
func newAccountsClient(client *Client) *AccountsClient {
	return &AccountsClient{
		pathUrl: "/accounts",
		client:  client,
	}
}

// GetOperationKey retrieves the encrypted operation key for the authenticated account.
// This key is required for transaction signing and must be decrypted using the operation passcode.
//
// Returns:
//   - *GetOperationKeyResponse: Contains the encrypted operation key and its hash
//   - error: nil on success, error on failure
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

// CreateNewAPIKey generates a new API key for the authenticated account.
//
// Returns:
//   - *CreateNewAPIKeyResponse: Contains the new API key
//   - error: nil on success, error on failure
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

// GetDepositRecords retrieves all deposit transaction records for the authenticated account.
//
// Returns:
//   - *GetDepositRecordsResponse: Array of deposit records with status, assets, and transaction hashes
//   - error: nil on success, error on failure
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

// GetWithdrawalRecords retrieves all withdrawal transaction records for the authenticated account.
//
// Returns:
//   - *GetWithdrawalRecordsResponse: Array of withdrawal records with status and assets
//   - error: nil on success, error on failure
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

// GetOrderRecords retrieves order records based on the specified status and pagination parameters.
// Supports filtering by status (open orders, order history, trading history), symbol, and pagination.
//
// Parameters:
//   - data: Request parameters including status, limit, page, and optional symbol filter
//
// Returns:
//   - *GetOrderRecordsResponse: Paginated order records with total count and page info
//   - error: nil on success, error on failure
func (c *AccountsClient) GetOrderRecords(data *GetOrderRecordRequest) (*GetOrderRecordsResponse, error) {
	// Build query parameters
	params := make(map[string]string)
	params["status"] = string(data.Status)

	if data.Limit > 0 {
		params["limit"] = strconv.Itoa(data.Limit)
	}

	if data.Page > 0 {
		params["page"] = strconv.Itoa(data.Page)
	}

	if data.Symbol != "" {
		params["symbol"] = string(data.Symbol)
	}

	// Get request with query parameters
	bodyBytes, err := c.client.getWithParams(c.pathUrl+"/order-records", params)
	if err != nil {
		return nil, err
	}

	var getOrderRecordsResponse GetOrderRecordsResponse
	err = json.Unmarshal(bodyBytes, &getOrderRecordsResponse)
	if err != nil {
		return nil, err
	}
	return &getOrderRecordsResponse, nil
}

// GetOrderRecord retrieves a single order record by order ID.
//
// Parameters:
//   - orderId: The unique identifier of the order
//
// Returns:
//   - *GetOrderRecordResponse: Complete order details
//   - error: nil on success, error on failure
func (c *AccountsClient) GetOrderRecord(orderId string) (*GetOrderRecordResponse, error) {
	// Get request with query parameters - note the endpoint is /account/order (singular)
	bodyBytes, err := c.client.get(c.pathUrl + "/order/" + orderId)
	if err != nil {
		return nil, err
	}

	var getOrderRecordResponse GetOrderRecordResponse
	err = json.Unmarshal(bodyBytes, &getOrderRecordResponse)

	if err != nil {
		fmt.Printf("Error unmarshalling GetOrderRecordResponse: %v\n", err)
		return nil, err
	}
	return &getOrderRecordResponse, nil
}

// GetAccountBalance retrieves the current balance for all assets in the authenticated account.
//
// Returns:
//   - *GetAccountBalanceResponse: Array of asset balances showing free and locked amounts
//   - error: nil on success, error on failure
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

// BuildDepositTransaction builds a deposit transaction for the specified assets.
// The returned transaction hex must be signed and then submitted using SubmitDepositTransaction.
//
// Parameters:
//   - data: Deposit request containing assets to deposit and input UTXOs
//
// Returns:
//   - *BuildDepositTransactionResponse: Transaction hex ready for signing
//   - error: nil on success, error on failure
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

// BuildWithdrawalTransaction builds a withdrawal transaction for the specified assets.
// The returned transaction hex must be signed and then submitted using SubmitWithdrawalTransaction.
//
// Parameters:
//   - data: Withdrawal request containing assets to withdraw
//
// Returns:
//   - *BuildWithdrawalTransactionResponse: Transaction hex ready for signing
//   - error: nil on success, error on failure
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

// BuildTransferalTransaction builds a transfer transaction to send assets to another address.
// The returned transaction hex must be signed and then submitted using SubmitTransferalTransaction.
//
// Parameters:
//   - data: Transfer request containing assets to transfer and destination address
//
// Returns:
//   - *BuildTransferalTransactionResponse: Transaction hex ready for signing
//   - error: nil on success, error on failure
func (c *AccountsClient) BuildTransferalTransaction(data *BuildTransferalTransactionRequest) (*BuildTransferalTransactionResponse, error) {
	bodyBytes, err := c.client.post(c.pathUrl+"/transferal/build", data)
	if err != nil {
		return nil, err
	}

	var buildTransferalTransactionResponse BuildTransferalTransactionResponse
	err = json.Unmarshal(bodyBytes, &buildTransferalTransactionResponse)
	if err != nil {
		return nil, err
	}
	return &buildTransferalTransactionResponse, nil
}

// SubmitDepositTransaction submits a signed deposit transaction to the network.
//
// Parameters:
//   - data: Submit request containing the signed transaction hex
//
// Returns:
//   - *SubmitDepositTransactionResponse: Transaction hash of the submitted transaction
//   - error: nil on success, error on failure
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

// SubmitWithdrawalTransaction submits a signed withdrawal transaction to the network.
//
// Parameters:
//   - data: Submit request containing the signed transaction hex
//
// Returns:
//   - *SubmitWithdrawalTransactionResponse: Transaction hash of the submitted transaction
//   - error: nil on success, error on failure
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

// SubmitTransferalTransaction submits a signed transfer transaction to the network.
//
// Parameters:
//   - data: Submit request containing the signed transaction hex
//
// Returns:
//   - *SubmitTransferalTransactionResponse: Transaction hash of the submitted transaction
//   - error: nil on success, error on failure
func (c *AccountsClient) SubmitTransferalTransaction(data *SubmitTransferalTransactionRequest) (*SubmitTransferalTransactionResponse, error) {
	bodyBytes, err := c.client.post(c.pathUrl+"/transferal/submit", data)
	if err != nil {
		return nil, err
	}

	var submitTransferalTransactionResponse SubmitTransferalTransactionResponse
	err = json.Unmarshal(bodyBytes, &submitTransferalTransactionResponse)
	if err != nil {
		return nil, err
	}
	return &submitTransferalTransactionResponse, nil
}
