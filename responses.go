package deltadefi

type GetOperationKeyResponse struct {
	EncryptedOperationKey string `json:"encrypted_operation_key"`
	OperationKeyHash      string `json:"operation_key_hash"`
}

type BuildSendRefScriptsTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type SubmitSendRefScriptsTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}
type SubmitPostOrderTransactionResponse struct {
	Order   *OrderJSON `json:"order"`
	TxHexes string     `json:"tx_hexes"`
}

type PostOrderResponse = *SubmitPostOrderTransactionResponse

type DepositRecord struct {
	CreatedAt string            `json:"created_at"`
	Status    TransactionStatus `json:"status"`
	Assets    []Asset           `json:"assets"`
	TxHash    string            `json:"tx_hash"`
}

type GetDepositRecordsResponse []DepositRecord

// GetOrderRecordResponse represents the response from the get order records endpoint.
type GetOrderRecordResponse struct {
	Orders              []OrderJSON              `json:"orders"`
	OrderFillingRecords []OrderFillingRecordJSON `json:"order_filling_records"`
}

type WithdrawalRecord struct {
	CreatedAt string            `json:"created_at"`
	Status    TransactionStatus `json:"status"`
	Assets    []Asset           `json:"assets"`
}

type GetWithdrawalRecordsResponse []WithdrawalRecord

type AssetBalance struct {
	Asset  string `json:"asset"`
	Free   int64  `json:"free"`
	Locked int64  `json:"locked"`
}

type GetAccountBalanceResponse []AssetBalance

type CreateNewAPIKeyResponse struct {
	APIKey string `json:"api_key"`
}

type BuildDepositTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type BuildWithdrawalTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type BuildTransferalTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type SubmitDepositTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

type SubmitWithdrawalTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

type SubmitTransferalTransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

type MarketDepth struct {
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
}

type GetMarketDepthResponse struct {
	Bids []MarketDepth `json:"bids"`
	Asks []MarketDepth `json:"asks"`
}

type GetMarketPriceResponse struct {
	Price float64 `json:"price"`
}

type Trade struct {
	Time   string  `json:"time"`
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
}

type GetAggregatedPriceResponse []Trade

type BuildPlaceOrderTransactionResponse struct {
	OrderID string `json:"order_id"`
	TxHex   string `json:"tx_hex"`
}

type BuildCancelOrderTransactionResponse struct {
	TxHex string `json:"tx_hex"`
}

type SubmitPlaceOrderTransactionResponse struct {
	Order OrderJSON `json:"order"`
}

type SubmitCancelOrderTransactionResponse struct {
	TxHash string `json:"txhash"`
}
